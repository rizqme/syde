package cli

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/audit"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var (
	addDescription string
	addPurpose     string
	addTags        []string
	addFiles       []string
	addBody        string
	addRels        []string
	addNotes       []string
	// Component
	addResponsibility  string
	addCapabilities    []string
	addBoundaries      string
	addBehaviorSummary string
	addInteractionSum  string
	addDataHandling    string
	addScalingNotes    string
	// Contract
	addContractKind       string
	addInteractionPattern string
	addProtocolNotes      string
	addInput              string
	addInputParams        []string
	addOutput             string
	addOutputParams       []string
	addConstraintsText    string
	addVersioningNotes    string
	addContractWireframe  string
	// Concept
	addMeaning    string
	addLifecycle  string
	addInvariants string
	// Flow
	addTrigger      string
	addGoal         string
	addNarrative    string
	addHappyPath    string
	addEdgeCases    string
	addFailureModes string
	addPerfNotes    string
	addFlowSteps    []string
	// Requirement (statement/rationale/supersedes shared with legacy decision flags)
	addStatement  string
	addRationale  string
	addSupersedes string
	// Requirement
	addReqType           string
	addReqPriority       string
	addReqVerification   string
	addReqSource         string
	addReqSourceRef      string
	addReqStatus         string
	addReqSupersededBy   string
	addReqObsoleteReason string
	addReqApprovedAt     string
	addReqAuditedOverlaps []string
	addReqForce           bool
	// System
	addSysContext     string
	addSysScope       string
	addSysPrinciples  string
	addSysQuality     string
	addSysAssumptions string
)

// parseContractParams converts "path|type|description" specs into ContractParam slices.
func parseContractParams(specs []string) []model.ContractParam {
	var out []model.ContractParam
	for _, s := range specs {
		if s == "" {
			continue
		}
		if p, ok := model.ParseContractParam(s); ok {
			out = append(out, p)
		}
	}
	return out
}

// parseFlowSteps converts "id|action|contract|description|on_success|on_failure"
// specs into FlowStep slices. Minimum 2 fields (id + action), rest optional.
func parseFlowSteps(specs []string) []model.FlowStep {
	var out []model.FlowStep
	for _, s := range specs {
		if s == "" {
			continue
		}
		parts := strings.SplitN(s, "|", 6)
		if len(parts) < 2 || strings.TrimSpace(parts[0]) == "" {
			continue
		}
		step := model.FlowStep{
			ID:     strings.TrimSpace(parts[0]),
			Action: strings.TrimSpace(parts[1]),
		}
		if len(parts) > 2 {
			step.Contract = strings.TrimSpace(parts[2])
		}
		if len(parts) > 3 {
			step.Description = strings.TrimSpace(parts[3])
		}
		if len(parts) > 4 {
			step.OnSuccess = strings.TrimSpace(parts[4])
		}
		if len(parts) > 5 {
			step.OnFailure = strings.TrimSpace(parts[5])
		}
		out = append(out, step)
	}
	return out
}

func validEntityKinds() string {
	kinds := model.AllEntityKinds()
	names := make([]string, 0, len(kinds))
	for _, kind := range kinds {
		names = append(names, string(kind))
	}
	return strings.Join(names, ", ")
}

func parseRefList(specs ...string) []string {
	var out []string
	for _, spec := range specs {
		for _, part := range strings.Split(spec, ",") {
			ref := strings.TrimSpace(part)
			if ref != "" {
				out = append(out, ref)
			}
		}
	}
	return out
}

func appendRefOnce(refs []string, ref string) []string {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return refs
	}
	for _, existing := range refs {
		if existing == ref {
			return refs
		}
	}
	return append(refs, ref)
}

func parseRequirementStatus(raw string) (model.RequirementStatus, error) {
	status := model.RequirementStatus(strings.TrimSpace(strings.ToLower(raw)))
	if status == "" {
		return model.RequirementActive, nil
	}
	switch status {
	case model.RequirementActive, model.RequirementSuperseded, model.RequirementObsolete:
		return status, nil
	default:
		return "", fmt.Errorf("invalid requirement status %q: expected active, superseded, or obsolete", raw)
	}
}

// parseRelSpec parses a --add-rel value as either "target:type" or
// "target:type:label". Label is used for cardinality on relates_to
// ("one-to-many" and friends) or any other free-form relationship
// annotation. Two-part form is preserved for back-compat with every
// existing entity on disk. Errors carry the original spec so the
// user can see exactly what was rejected.
func parseRelSpec(spec string) (model.Relationship, error) {
	parts := strings.SplitN(spec, ":", 3)
	if len(parts) < 2 {
		return model.Relationship{}, fmt.Errorf("invalid --add-rel %q: expected target:type or target:type:label", spec)
	}
	r := model.Relationship{
		Target: strings.TrimSpace(parts[0]),
		Type:   strings.TrimSpace(parts[1]),
	}
	if r.Target == "" || r.Type == "" {
		return model.Relationship{}, fmt.Errorf("invalid --add-rel %q: target and type must be non-empty", spec)
	}
	if len(parts) == 3 {
		r.Label = strings.TrimSpace(parts[2])
	}
	return r, nil
}

var addCmd = &cobra.Command{
	Use:   "add <kind> <name>",
	Short: "Add a new entity to the design model",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		kindStr := strings.ToLower(args[0])
		name := args[1]

		kind, ok := model.ParseEntityKind(kindStr)
		if !ok {
			return fmt.Errorf("unknown entity kind: %s\nValid: %s", kindStr, validEntityKinds())
		}

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity := model.NewEntityForKind(kind)
		b := entity.GetBase()
		b.Name = name
		b.Description = addDescription
		b.Purpose = addPurpose
		b.Tags = addTags
		b.Files = addFiles
		b.Notes = addNotes
		for _, spec := range addRels {
			if spec == "" {
				continue
			}
			rel, err := parseRelSpec(spec)
			if err != nil {
				return err
			}
			b.Relationships = append(b.Relationships, rel)
		}

		switch v := entity.(type) {
		case *model.ComponentEntity:
			v.Responsibility = addResponsibility
			v.Capabilities = addCapabilities
			v.Boundaries = addBoundaries
			v.BehaviorSummary = addBehaviorSummary
			v.InteractionSummary = addInteractionSum
			v.DataHandling = addDataHandling
			v.ScalingNotes = addScalingNotes
		case *model.ContractEntity:
			v.ContractKind = addContractKind
			v.InteractionPattern = addInteractionPattern
			v.ProtocolNotes = addProtocolNotes
			v.Input = addInput
			v.InputParameters = parseContractParams(addInputParams)
			v.Output = addOutput
			v.OutputParameters = parseContractParams(addOutputParams)
			v.Constraints = addConstraintsText
			v.VersioningNotes = addVersioningNotes
			v.Wireframe = addContractWireframe
		case *model.ConceptEntity:
			v.Meaning = addMeaning
			v.Lifecycle = addLifecycle
			v.Invariants = addInvariants
		case *model.FlowEntity:
			v.Trigger = addTrigger
			v.Goal = addGoal
			v.Narrative = addNarrative
			v.HappyPath = addHappyPath
			v.EdgeCases = addEdgeCases
			v.FlowFailureModes = addFailureModes
			v.PerformanceNotes = addPerfNotes
			v.Steps = parseFlowSteps(addFlowSteps)
		case *model.RequirementEntity:
			status, err := parseRequirementStatus(addReqStatus)
			if err != nil {
				return err
			}
			v.Statement = addStatement
			v.ReqType = model.RequirementType(addReqType)
			v.Priority = model.RequirementPriority(addReqPriority)
			v.Verification = addReqVerification
			v.Source = addReqSource
			if v.Source == "" {
				v.Source = "manual"
			}
			v.SourceRef = addReqSourceRef
			v.RequirementStatus = status
			v.Rationale = addRationale
			v.Supersedes = parseRefList(addSupersedes)
			v.SupersededBy = parseRefList(addReqSupersededBy)
			v.ObsoleteReason = addReqObsoleteReason
			v.ApprovedAt = addReqApprovedAt
			v.AuditedOverlaps = parseAuditedOverlaps(addReqAuditedOverlaps)
			if !hasRelationshipType(b.Relationships, model.RelBelongsTo) {
				if parent := rootSystemTarget(store); parent != "" {
					b.Relationships = append(b.Relationships, model.Relationship{Target: parent, Type: model.RelBelongsTo})
				}
			}
		case *model.SystemEntity:
			v.Context = addSysContext
			v.Scope = addSysScope
			v.DesignPrinciples = addSysPrinciples
			v.QualityGoals = addSysQuality
			v.Assumptions = addSysAssumptions
		}

		// Pre-create validation: skip the ID check because Create() is
		// what allocates the ID. Any remaining errors are real and
		// should block the write.
		var contentErrs []model.ValidationError
		for _, ve := range model.ValidateEntity(entity) {
			if ve.Field == "id" {
				continue
			}
			contentErrs = append(contentErrs, ve)
		}
		if len(contentErrs) > 0 {
			var msgs []string
			for _, ve := range contentErrs {
				msgs = append(msgs, fmt.Sprintf("  - %s: %s", ve.Field, ve.Message))
			}
			return fmt.Errorf("validation failed:\n%s", strings.Join(msgs, "\n"))
		}

		// Pre-creation overlap detection for requirements using TF-IDF.
		// Blocks the create unless every overlap above the threshold is
		// acknowledged via --audited slug[:reason] or the author passes
		// --force. This shifts the overlap gate to the earliest possible
		// moment so no overlapping entity is written to disk.
		if req, ok := entity.(*model.RequirementEntity); ok && req.Statement != "" {
			overlapMsgs, unackedSlugs, err := detectRequirementOverlaps(store, req, b.ID, parseAuditedOverlaps(addReqAuditedOverlaps))
			if err != nil {
				return fmt.Errorf("overlap detection: %w", err)
			}
			if len(overlapMsgs) > 0 {
				fmt.Println()
				fmt.Println("⚠ overlap candidates detected — semantic review required:")
				for _, o := range overlapMsgs {
					fmt.Println(o)
				}
				if len(unackedSlugs) > 0 && !addReqForce {
					fmt.Println()
					fmt.Println("  Resolve each unacknowledged overlap before creating:")
					fmt.Println("    MERGE   — reuse the existing requirement, do not create a duplicate")
					fmt.Println("    RENAME  — rewrite this statement so the two read distinctly, then retry")
					fmt.Println("    DISTINCT— pass --audited <slug>:\"semantic distinction rationale\" for each slug below")
					fmt.Println()
					for _, s := range unackedSlugs {
						fmt.Printf("      --audited %s:\"<why this requirement means something different>\"\n", s)
					}
					fmt.Println()
					fmt.Println("  Override with --force if the overlap is genuinely unavoidable (rare).")
					return fmt.Errorf("requirement creation blocked: %d unacknowledged overlap(s)", len(unackedSlugs))
				}
			}
		}

		filePath, err := store.Create(entity, addBody)
		if err != nil {
			return fmt.Errorf("create entity: %w", err)
		}

		fmt.Printf("Created %s: %s\n", kind, name)
		fmt.Printf("  ID: %s\n", b.ID)
		fmt.Printf("  File: %s\n", filePath)

		return nil
	},
}

// detectRequirementOverlaps runs TF-IDF similarity for a candidate
// requirement against all active requirements, returning human-readable
// overlap messages and the set of overlap slugs that are not yet
// acknowledged. selfID is excluded from the comparison.
func detectRequirementOverlaps(store *writeClient, candidate *model.RequirementEntity, selfID string, audited []model.AuditedOverlap) ([]string, []string, error) {
	newTerms := audit.SignificantTerms(candidate.Statement)
	if len(newTerms) == 0 {
		return nil, nil, nil
	}
	allEntities, err := store.List(model.KindRequirement)
	if err != nil {
		return nil, nil, err
	}
	type reqRef struct {
		id, name, slug string
		terms          map[string]bool
	}
	var allTermSets []map[string]bool
	var others []reqRef
	for _, ewb := range allEntities {
		other, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || other.RequirementStatus != model.RequirementActive || other.Statement == "" {
			continue
		}
		ob := other.GetBase()
		terms := audit.SignificantTerms(other.Statement)
		if len(terms) == 0 {
			continue
		}
		allTermSets = append(allTermSets, terms)
		if ob.ID != selfID {
			others = append(others, reqRef{id: ob.ID, name: ob.Name, slug: ob.CanonicalSlug(), terms: terms})
		}
	}
	corpus := audit.NewTFIDFCorpus(allTermSets)
	auditedSet := map[string]bool{}
	for _, a := range audited {
		auditedSet[a.Slug] = true
		auditedSet[utils.BaseSlug(a.Slug)] = true
	}
	var msgs []string
	var unacked []string
	for _, o := range others {
		sim := corpus.TFIDFSimilarity(newTerms, o.terms)
		if sim <= 0.6 {
			continue
		}
		msgs = append(msgs, fmt.Sprintf("  ⚠ %s %q (%.0f%% TF-IDF) — slug: %s", o.id, o.name, sim*100, o.slug))
		if !auditedSet[o.slug] && !auditedSet[utils.BaseSlug(o.slug)] {
			unacked = append(unacked, o.slug)
		}
	}
	return msgs, unacked, nil
}

func init() {
	f := addCmd.Flags()
	// Base
	f.StringVar(&addDescription, "description", "", "entity description")
	f.StringVar(&addPurpose, "purpose", "", "why the entity exists")
	f.StringSliceVar(&addTags, "tag", nil, "entity tags")
	// File / note / capability / rel use StringArrayVar (literal-valued,
	// repeatable) NOT StringSliceVar — slice splits on commas inside the
	// value, which corrupts paths, sentences with commas, and relationship
	// targets.
	f.StringArrayVar(&addFiles, "file", nil, "concrete source file path (repeatable)")
	f.StringArrayVar(&addNotes, "note", nil, "informal note (repeatable)")
	f.StringVar(&addBody, "body", "", "markdown body content")
	f.StringArrayVar(&addRels, "add-rel", nil, "add relationship (target:type) — repeatable")
	// Component
	f.StringVar(&addResponsibility, "responsibility", "", "component responsibility")
	f.StringArrayVar(&addCapabilities, "capability", nil, "component capability (repeatable)")
	f.StringVar(&addBoundaries, "boundaries", "", "component boundaries")
	f.StringVar(&addBehaviorSummary, "behavior-summary", "", "behavior summary")
	f.StringVar(&addInteractionSum, "interaction-summary", "", "interaction summary")
	f.StringVar(&addDataHandling, "data-handling", "", "data handling")
	f.StringVar(&addScalingNotes, "scaling-notes", "", "scaling notes")
	// Contract
	f.StringVar(&addContractKind, "contract-kind", "", "contract kind (api/event/command/query)")
	f.StringVar(&addInteractionPattern, "interaction-pattern", "", "interaction pattern")
	f.StringVar(&addProtocolNotes, "protocol-notes", "", "protocol notes")
	f.StringVar(&addInput, "input", "", "contract invocation signature (e.g. 'GET /api/projects', 'syde plan create <name>')")
	f.StringArrayVar(&addInputParams, "input-parameter", nil, "input parameter 'path|type|description' (repeatable)")
	f.StringVar(&addOutput, "output", "", "contract output signature / response shape")
	f.StringArrayVar(&addOutputParams, "output-parameter", nil, "output parameter 'path|type|description' (repeatable)")
	f.StringVar(&addConstraintsText, "constraints-text", "", "constraints")
	f.StringVar(&addVersioningNotes, "versioning-notes", "", "versioning notes")
	f.StringVar(&addContractWireframe, "wireframe", "", "screen contract UIML wireframe source — required when --contract-kind=screen")
	// Concept
	f.StringVar(&addMeaning, "meaning", "", "concept meaning (one-liner explaining what this domain term is)")
	f.StringVar(&addLifecycle, "lifecycle", "", "lifecycle prose for the concept")
	f.StringVar(&addInvariants, "invariants", "", "invariants — rules that must always hold")
	// Flow
	f.StringVar(&addTrigger, "trigger", "", "flow trigger")
	f.StringVar(&addGoal, "goal", "", "flow goal")
	f.StringVar(&addNarrative, "narrative", "", "flow narrative")
	f.StringVar(&addHappyPath, "happy-path", "", "happy path")
	f.StringVar(&addEdgeCases, "edge-cases", "", "edge cases")
	f.StringVar(&addFailureModes, "failure-modes", "", "failure modes")
	f.StringArrayVar(&addFlowSteps, "step", nil, "flow step 'id|action|contract|description|on_success|on_failure' (repeatable)")
	f.StringVar(&addPerfNotes, "performance-notes", "", "performance notes")
	// Requirement
	f.StringVar(&addStatement, "statement", "", "requirement statement (EARS shall-form)")
	f.StringVar(&addRationale, "rationale", "", "requirement rationale")
	f.StringVar(&addSupersedes, "supersedes", "", "slug/ref superseded by this requirement")
	f.StringVar(&addReqSource, "source", "", "requirement source (user/plan/migration/manual)")
	f.StringVar(&addReqSourceRef, "source-ref", "", "requirement source reference")
	f.StringVar(&addReqStatus, "requirement-status", "", "requirement status (active/superseded/obsolete)")
	f.StringVar(&addReqType, "type", "", "requirement type (functional/non-functional/constraint/interface/performance/security/usability)")
	f.StringVar(&addReqPriority, "priority", "", "requirement priority (must/should/could/wont)")
	f.StringVar(&addReqVerification, "verification", "", "how the requirement is verified (automated/integration-test/inspection/manual)")
	f.StringVar(&addReqSupersededBy, "superseded-by", "", "requirement refs that supersede this one, comma-separated")
	f.StringVar(&addReqObsoleteReason, "obsolete-reason", "", "why the requirement is obsolete")
	f.StringVar(&addReqApprovedAt, "approved-at", "", "requirement approval timestamp")
	f.StringArrayVar(&addReqAuditedOverlaps, "audited", nil, "acknowledged overlapping requirement as slug[:distinction rationale] (repeatable)")
	f.BoolVar(&addReqForce, "force", false, "bypass the overlap gate even when unacknowledged overlaps exist (rare)")
	// System
	f.StringVar(&addSysContext, "context-text", "", "system context")
	f.StringVar(&addSysScope, "scope", "", "system scope")
	f.StringVar(&addSysPrinciples, "design-principles", "", "design principles")
	f.StringVar(&addSysQuality, "quality-goals", "", "quality goals")
	f.StringVar(&addSysAssumptions, "assumptions", "", "assumptions")

	rootCmd.AddCommand(addCmd)
}
