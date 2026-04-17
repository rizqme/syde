package audit

import (
	"fmt"
	"math"
	"strings"

	"github.com/feedloop/syde/internal/model"
)

func planAuthoringFindings(all []model.EntityWithBody) []Finding {
	lookup, _ := auditGraph(all)
	taskNames := map[string]string{}
	for _, ewb := range all {
		if task, ok := ewb.Entity.(*model.TaskEntity); ok {
			b := task.GetBase()
			taskNames[b.CanonicalSlug()] = b.Name
		}
	}

	var out []Finding
	for _, ewb := range all {
		plan, ok := ewb.Entity.(*model.PlanEntity)
		if !ok || plan.PlanStatus == model.PlanCompleted {
			continue
		}
		b := plan.GetBase()
		planSlug := b.CanonicalSlug()
		planTasks := map[string]bool{}
		for _, slug := range plan.AllTasks() {
			planTasks[slug] = true
		}
		claimedTasks := map[string]bool{}

		if laneEntryCount(plan.Changes.Requirements) == 0 {
			out = append(out, planAuthoringFinding(SeverityWarning, b, "changes.requirements", "requirements lane is empty — the plan declares no durable property"))
		}

		lanes := []struct {
			name string
			kind model.EntityKind
			lane model.ChangeLane
		}{
			{"requirements", model.KindRequirement, plan.Changes.Requirements},
			{"systems", model.KindSystem, plan.Changes.Systems},
			{"concepts", model.KindConcept, plan.Changes.Concepts},
			{"components", model.KindComponent, plan.Changes.Components},
			{"contracts", model.KindContract, plan.Changes.Contracts},
			{"flows", model.KindFlow, plan.Changes.Flows},
		}

		for _, lane := range lanes {
			for _, d := range lane.lane.Deleted {
				out = append(out, validateChangeTasks(b, planTasks, d.Tasks, fmt.Sprintf("changes.%s.deleted[%s].tasks", lane.name, d.ID), d.Slug)...)
				for _, t := range d.Tasks {
					claimedTasks[t] = true
				}
			}
			for _, e := range lane.lane.Extended {
				// Empty field_changes is a legitimate hand-review
				// declaration; the completion validator only checks
				// declared field_changes, so nothing programmatic is
				// lost when the list is empty. The previous WARN
				// here was a documentation prompt and is dropped
				// under the strict-severity model.
				out = append(out, validateChangeTasks(b, planTasks, e.Tasks, fmt.Sprintf("changes.%s.extended[%s].tasks", lane.name, e.ID), e.Slug)...)
				for _, t := range e.Tasks {
					claimedTasks[t] = true
				}
				if lane.kind == model.KindContract {
					if target, ok := lookup[e.Slug]; ok {
						if contract, ok := target.Entity.(*model.ContractEntity); ok && contract.ContractKind == "screen" {
							if _, ok := e.FieldChanges["wireframe"]; !ok {
								out = append(out, planAuthoringFinding(
									SeverityWarning,
									b,
									fmt.Sprintf("changes.%s.extended[%s].field_changes.wireframe", lane.name, e.ID),
									fmt.Sprintf("screen contract %s is extended without a wireframe field_change", e.Slug),
								))
							}
						}
					}
				}
			}
			for _, n := range lane.lane.New {
				out = append(out, validateNewDraft(b, lane.kind, lane.name, n)...)
				out = append(out, validateChangeTasks(b, planTasks, n.Tasks, fmt.Sprintf("changes.%s.new[%s].tasks", lane.name, n.ID), n.Name)...)
				for _, t := range n.Tasks {
					claimedTasks[t] = true
				}
			}
		}

		// Requirement coverage: requirements should be proportional to
		// non-requirement changes. A plan with many implementation changes
		// and few requirements is likely under-specified.
		reqCount := laneEntryCount(plan.Changes.Requirements)
		nonReqCount := 0
		for _, lane := range lanes {
			if lane.kind == model.KindRequirement {
				continue
			}
			nonReqCount += laneEntryCount(lane.lane)
		}
		if nonReqCount > 0 && reqCount*3 < nonReqCount {
			out = append(out, planAuthoringFinding(
				SeverityWarning,
				b,
				"changes.requirements",
				fmt.Sprintf("plan has %d requirement(s) for %d non-requirement change(s) — consider decomposing into more granular requirements (aim for at least 1 requirement per 3 changes)", reqCount, nonReqCount),
			))
		}

		// Requirement overlap: for each new requirement in the plan,
		// check if an existing active requirement shares >50% of
		// significant terms in its statement.
		for _, n := range plan.Changes.Requirements.New {
			statement := draftString(n.Draft, "statement")
			if statement == "" {
				continue
			}
			newTerms := SignificantTerms(statement)
			if len(newTerms) == 0 {
				continue
			}
			for _, ewb2 := range all {
				req, ok := ewb2.Entity.(*model.RequirementEntity)
				if !ok || req.RequirementStatus != model.RequirementActive {
					continue
				}
				if req.Statement == "" {
					continue
				}
				// Skip the entity that was just created from this
				// declared change (matched by name) — it overlaps
				// itself at 100% which is not a finding.
				if req.GetBase().Name == n.Name {
					continue
				}
				existingTerms := SignificantTerms(req.Statement)
				if len(existingTerms) == 0 {
					continue
				}
				overlap := TermOverlap(newTerms, existingTerms)
				if overlap > 0.5 {
					out = append(out, planAuthoringFinding(
						SeverityWarning,
						b,
						fmt.Sprintf("changes.requirements.new[%s]", n.ID),
						fmt.Sprintf("new requirement %q may overlap existing %q (%.0f%% term similarity) — link via refines/derives_from or supersede the existing one", n.Name, req.GetBase().Name, overlap*100),
					))
				}
			}
		}

		for taskSlug := range planTasks {
			if claimedTasks[taskSlug] || isVerificationTask(taskNames[taskSlug], taskSlug) {
				continue
			}
			out = append(out, Finding{
				Severity:   SeverityWarning,
				Category:   CatPlanAuthoring,
				Message:    fmt.Sprintf("task %q is not referenced by any change.tasks list", taskSlug),
				EntityKind: model.KindPlan,
				EntitySlug: planSlug,
				EntityName: b.Name,
				Field:      "phases.tasks",
			})
		}

		out = append(out, contractCoverageFindings(b, planSlug, plan, lookup)...)
		out = append(out, flowCoverageFindings(b, planSlug, plan, lookup)...)
	}
	return out
}

// contractCoverageFindings emits a warning for every requirement
// surface (CLI/REST/screen/event) named in a new/extended requirement
// change that is NOT covered by any contract (new or extended) in the
// same plan diff. Surfaces are extracted via ExtractSurfaces; coverage
// against the plan's contract lane uses the contract's declared
// invocation (new entries' draft `input` field; extended entries'
// target contract's input plus any field_changes override).
func contractCoverageFindings(b *model.BaseEntity, planSlug string, plan *model.PlanEntity, lookup map[string]auditEntity) []Finding {
	// Gather the contract inputs this plan introduces or extends.
	var planContractInputs []string
	for _, n := range plan.Changes.Contracts.New {
		if input := draftString(n.Draft, "input"); input != "" {
			planContractInputs = append(planContractInputs, input)
		}
		// Also allow the contract's declared name to match the surface
		// (e.g. a contract named "Place Order" for a requirement
		// mentioning "place order" as a surface).
		if n.Name != "" {
			planContractInputs = append(planContractInputs, n.Name)
		}
	}
	for _, e := range plan.Changes.Contracts.Extended {
		if override, ok := e.FieldChanges["input"]; ok {
			planContractInputs = append(planContractInputs, fmt.Sprintf("%v", override))
		}
		// Fall back to the current contract input.
		if target, ok := lookup[e.Slug]; ok {
			if contract, ok := target.Entity.(*model.ContractEntity); ok {
				planContractInputs = append(planContractInputs, contract.Input)
			}
		}
	}
	// Also accept any active contract in the model — the planning
	// rule should only fire when a surface is genuinely uncovered,
	// not when the plan happens to mention an existing covered
	// surface. Mirrors the post-plan rule's broader scope.
	for _, ent := range lookup {
		contract, ok := ent.Entity.(*model.ContractEntity)
		if !ok {
			continue
		}
		if contract.Input != "" {
			planContractInputs = append(planContractInputs, contract.Input)
		}
		if contract.Name != "" {
			planContractInputs = append(planContractInputs, contract.Name)
		}
	}

	check := func(stmt, field, humanName string) []Finding {
		var out []Finding
		for _, surface := range ExtractSurfaces(stmt) {
			covered := false
			for _, input := range planContractInputs {
				if ContractCoversSurface(input, surface) {
					covered = true
					break
				}
			}
			if covered {
				continue
			}
			out = append(out, planAuthoringFinding(
				SeverityWarning,
				b,
				field,
				fmt.Sprintf("requirement %q names %s surface %q but no contract in this plan's diff covers it — add or extend a contract whose input matches, or reword the requirement", humanName, surface.Kind, surface.Raw),
			))
		}
		return out
	}

	var findings []Finding
	for _, n := range plan.Changes.Requirements.New {
		if stmt := draftString(n.Draft, "statement"); stmt != "" {
			findings = append(findings, check(stmt, fmt.Sprintf("changes.requirements.new[%s].surfaces", n.ID), n.Name)...)
		}
	}
	for _, e := range plan.Changes.Requirements.Extended {
		stmt := ""
		if override, ok := e.FieldChanges["statement"]; ok {
			stmt = fmt.Sprintf("%v", override)
		}
		if stmt == "" {
			if target, ok := lookup[e.Slug]; ok {
				if req, ok := target.Entity.(*model.RequirementEntity); ok {
					stmt = req.Statement
				}
			}
		}
		if stmt != "" {
			findings = append(findings, check(stmt, fmt.Sprintf("changes.requirements.extended[%s].surfaces", e.ID), e.Slug)...)
		}
	}
	return findings
}

// flowCoverageFindings emits a warning for every contract introduced
// or extended by this plan that is not referenced by at least one
// flow in the plan's flow lane. Mirrors contractFlowFindings at
// planning time so the co-evolution gap is caught before the plan is
// approved.
func flowCoverageFindings(b *model.BaseEntity, planSlug string, plan *model.PlanEntity, lookup map[string]auditEntity) []Finding {
	// Set of contract slugs that at least one flow in the plan
	// references via its steps.
	referenced := map[string]bool{}
	collectSteps := func(stepsVal interface{}) {
		items, ok := stepsVal.([]interface{})
		if !ok {
			return
		}
		for _, item := range items {
			m, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			if c, ok := m["contract"].(string); ok && c != "" {
				referenced[c] = true
			}
		}
	}
	for _, n := range plan.Changes.Flows.New {
		collectSteps(n.Draft["steps"])
	}
	for _, e := range plan.Changes.Flows.Extended {
		if override, ok := e.FieldChanges["steps"]; ok {
			collectSteps(override)
		}
		// Fall back: if a flow is extended without a steps override,
		// consider every step on the current flow entity as already
		// referencing its contracts.
		if target, ok := lookup[e.Slug]; ok {
			if flow, ok := target.Entity.(*model.FlowEntity); ok {
				for _, step := range flow.Steps {
					if step.Contract != "" {
						referenced[step.Contract] = true
					}
				}
			}
		}
	}
	// Also count active flows in the broader model — the planning
	// rule mirrors the post-plan contract-flow coverage rule and
	// should not bark on a contract that is globally covered.
	for _, ent := range lookup {
		flow, ok := ent.Entity.(*model.FlowEntity)
		if !ok {
			continue
		}
		for _, step := range flow.Steps {
			if step.Contract != "" {
				referenced[step.Contract] = true
			}
		}
	}

	var findings []Finding
	for _, n := range plan.Changes.Contracts.New {
		// A new contract's slug isn't known yet; match against the
		// declared name slugified. Flows often reference the bare name.
		if !referenced[slugifyName(n.Name)] && !referenced[n.Name] {
			findings = append(findings, planAuthoringFinding(
				SeverityWarning,
				b,
				fmt.Sprintf("changes.contracts.new[%s].flow_coverage", n.ID),
				fmt.Sprintf("contract %q is introduced by this plan but no flow in the plan's flow lane references it — add a new flow or extend an existing flow's steps to include the contract", n.Name),
			))
		}
	}
	for _, e := range plan.Changes.Contracts.Extended {
		if !referenced[e.Slug] && !referenced[slugifyName(e.Slug)] {
			findings = append(findings, planAuthoringFinding(
				SeverityWarning,
				b,
				fmt.Sprintf("changes.contracts.extended[%s].flow_coverage", e.ID),
				fmt.Sprintf("contract %q is extended by this plan but no flow in the plan's flow lane references it — add a new flow or extend an existing flow's steps to include the contract", e.Slug),
			))
		}
	}
	return findings
}

// slugifyName collapses a human-readable contract name to a
// best-effort slug that matches how flows typically reference it.
func slugifyName(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

func laneEntryCount(lane model.ChangeLane) int {
	return len(lane.Deleted) + len(lane.Extended) + len(lane.New)
}

func planAuthoringFinding(sev Severity, b *model.BaseEntity, field, msg string) Finding {
	return Finding{
		Severity:   sev,
		Category:   CatPlanAuthoring,
		Message:    msg,
		EntityKind: model.KindPlan,
		EntitySlug: b.CanonicalSlug(),
		EntityName: b.Name,
		Field:      field,
	}
}

func validateChangeTasks(b *model.BaseEntity, planTasks map[string]bool, tasks []string, field, target string) []Finding {
	if len(tasks) == 0 {
		return []Finding{planAuthoringFinding(SeverityError, b, field, fmt.Sprintf("change %q has no implementing task slugs", target))}
	}
	var out []Finding
	for _, taskSlug := range tasks {
		if !planTasks[taskSlug] {
			out = append(out, planAuthoringFinding(SeverityError, b, field, fmt.Sprintf("change %q references task %q, but that task is not in this plan", target, taskSlug)))
		}
	}
	return out
}

func validateNewDraft(b *model.BaseEntity, kind model.EntityKind, laneName string, n model.NewChange) []Finding {
	field := fmt.Sprintf("changes.%s.new[%s].draft", laneName, n.ID)
	missing := func(msg string) Finding {
		return planAuthoringFinding(SeverityError, b, field, fmt.Sprintf("new %s %q draft %s", kind, n.Name, msg))
	}
	switch kind {
	case model.KindComponent:
		if strings.TrimSpace(draftString(n.Draft, "responsibility")) == "" {
			return []Finding{missing("must declare responsibility")}
		}
	case model.KindContract:
		var out []Finding
		for _, key := range []string{"contract_kind", "input", "output"} {
			if strings.TrimSpace(draftString(n.Draft, key)) == "" {
				out = append(out, missing("must declare "+key))
			}
		}
		return out
	case model.KindRequirement:
		statement := draftString(n.Draft, "statement")
		if _, ok := model.MatchEARS(statement); !ok {
			return []Finding{missing("must declare an EARS-compliant statement")}
		}
	case model.KindConcept:
		var out []Finding
		for _, key := range []string{"meaning", "invariants"} {
			if strings.TrimSpace(draftString(n.Draft, key)) == "" {
				out = append(out, missing("must declare "+key))
			}
		}
		return out
	}
	return nil
}

func draftString(draft map[string]interface{}, key string) string {
	if draft == nil {
		return ""
	}
	v, ok := draft[key]
	if !ok || v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

func isVerificationTask(name, slug string) bool {
	text := strings.ToLower(name + " " + slug)
	for _, token := range []string{"build", "test", "smoke", "refresh-tree", "verify", "verification", "validator"} {
		if strings.Contains(text, token) {
			return true
		}
	}
	return false
}

// stopWords are common English words excluded from term similarity.
var stopWords = map[string]bool{
	"the": true, "a": true, "an": true, "and": true, "or": true,
	"is": true, "are": true, "was": true, "were": true, "be": true,
	"been": true, "being": true, "have": true, "has": true, "had": true,
	"do": true, "does": true, "did": true, "will": true, "would": true,
	"shall": true, "should": true, "may": true, "might": true, "must": true,
	"can": true, "could": true, "to": true, "of": true, "in": true,
	"for": true, "on": true, "with": true, "at": true, "by": true,
	"from": true, "as": true, "into": true, "that": true, "which": true,
	"when": true, "where": true, "while": true, "if": true, "then": true,
	"than": true, "but": true, "not": true, "no": true, "nor": true,
	"its": true, "it": true, "this": true, "each": true,
	"any": true, "all": true, "both": true, "more": true, "most": true,
	"other": true, "some": true, "such": true, "only": true, "own": true,
	"same": true, "so": true, "also": true, "very": true, "just": true,
}

// SignificantTerms extracts lowercase words >3 chars, excluding stop words.
func SignificantTerms(text string) map[string]bool {
	terms := map[string]bool{}
	for _, word := range strings.Fields(strings.ToLower(text)) {
		word = strings.Trim(word, ".,;:!?\"'()-/")
		if len(word) <= 3 || stopWords[word] {
			continue
		}
		terms[word] = true
	}
	return terms
}

// TermOverlap is kept for the plan-authoring WARN (draft statements
// don't participate in a TF-IDF corpus). For entity-level overlap
// detection, use TFIDFSimilarity instead.
func TermOverlap(a, b map[string]bool) float64 {
	smaller, larger := a, b
	if len(a) > len(b) {
		smaller, larger = b, a
	}
	if len(smaller) == 0 {
		return 0
	}
	shared := 0
	for term := range smaller {
		if larger[term] {
			shared++
		}
	}
	return float64(shared) / float64(len(smaller))
}

// TFIDFCorpus builds inverse document frequency weights from a set of
// documents. Each document is a set of significant terms (use
// SignificantTerms to produce them). The corpus is then used by
// TFIDFSimilarity to compute cosine similarity between any two
// documents, with common terms (e.g. "syde", "entity", "audit")
// naturally down-weighted.
type TFIDFCorpus struct {
	docCount int
	df       map[string]int // how many documents contain each term
}

// NewTFIDFCorpus builds a corpus from a slice of term sets.
func NewTFIDFCorpus(docs []map[string]bool) *TFIDFCorpus {
	df := make(map[string]int)
	for _, doc := range docs {
		for term := range doc {
			df[term]++
		}
	}
	return &TFIDFCorpus{docCount: len(docs), df: df}
}

// idf returns the inverse document frequency for a term.
// Uses log(N/df) with a floor of 0 for terms in every document.
func (c *TFIDFCorpus) idf(term string) float64 {
	df := c.df[term]
	if df == 0 || c.docCount == 0 {
		return 0
	}
	return math.Log(float64(c.docCount) / float64(df))
}

// TFIDFSimilarity computes cosine similarity between two documents
// using TF-IDF weights from this corpus. Returns 0.0-1.0.
func (c *TFIDFCorpus) TFIDFSimilarity(a, b map[string]bool) float64 {
	// Build the union of terms
	allTerms := make(map[string]bool)
	for t := range a {
		allTerms[t] = true
	}
	for t := range b {
		allTerms[t] = true
	}
	if len(allTerms) == 0 {
		return 0
	}

	// TF is binary (1 if present, 0 if not) since we use term sets.
	// Weight = tf * idf = idf when present, 0 when absent.
	var dot, magA, magB float64
	for term := range allTerms {
		idf := c.idf(term)
		wa, wb := 0.0, 0.0
		if a[term] {
			wa = idf
		}
		if b[term] {
			wb = idf
		}
		dot += wa * wb
		magA += wa * wa
		magB += wb * wb
	}
	if magA == 0 || magB == 0 {
		return 0
	}
	return dot / (math.Sqrt(magA) * math.Sqrt(magB))
}
