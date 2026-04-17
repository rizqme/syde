package audit

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// reqEntry is a small bundle of an active requirement with its
// pre-tokenised statement. Shared by the overlap-detection and
// rubber-stamp rules.
type reqEntry struct {
	req   *model.RequirementEntity
	terms map[string]bool
}

func requirementFindings(all []model.EntityWithBody) []Finding {
	reqs := map[string]*model.RequirementEntity{}
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok {
			continue
		}
		addRequirementRef(reqs, req.ID, req)
		addRequirementRef(reqs, req.CanonicalSlug(), req)
		addRequirementRef(reqs, utils.BaseSlug(req.CanonicalSlug()), req)
		addRequirementRef(reqs, utils.Slugify(req.Name), req)
	}

	var out []Finding
	for _, req := range uniqueRequirements(reqs) {
		for _, ref := range req.Supersedes {
			oldReq := reqs[ref]
			if oldReq == nil {
				out = append(out, requirementFinding(req, "supersedes", fmt.Sprintf("supersedes target %q is not a requirement", ref)))
				continue
			}
			if oldReq.RequirementStatus != model.RequirementSuperseded {
				out = append(out, requirementFinding(oldReq, "requirement_status", fmt.Sprintf("must be superseded because %s supersedes it", req.CanonicalSlug())))
			}
			if !requirementRefs(oldReq.SupersededBy, req) {
				out = append(out, requirementFinding(oldReq, "superseded_by", fmt.Sprintf("must reference superseding requirement %s", req.CanonicalSlug())))
			}
		}
		for _, ref := range req.SupersededBy {
			newReq := reqs[ref]
			if newReq == nil {
				out = append(out, requirementFinding(req, "superseded_by", fmt.Sprintf("superseded_by target %q is not a requirement", ref)))
				continue
			}
			if req.RequirementStatus != model.RequirementSuperseded {
				out = append(out, requirementFinding(req, "requirement_status", "must be superseded when superseded_by is set"))
			}
			if !requirementRefs(newReq.Supersedes, req) {
				out = append(out, requirementFinding(newReq, "supersedes", fmt.Sprintf("must reference superseded requirement %s", req.CanonicalSlug())))
			}
		}
	}
	return out
}

func addRequirementRef(reqs map[string]*model.RequirementEntity, ref string, req *model.RequirementEntity) {
	if ref != "" {
		reqs[ref] = req
	}
}

func uniqueRequirements(reqs map[string]*model.RequirementEntity) []*model.RequirementEntity {
	seen := map[*model.RequirementEntity]bool{}
	var out []*model.RequirementEntity
	for _, req := range reqs {
		if seen[req] {
			continue
		}
		seen[req] = true
		out = append(out, req)
	}
	return out
}

func requirementRefs(refs []string, req *model.RequirementEntity) bool {
	for _, ref := range refs {
		if ref == req.ID || ref == req.CanonicalSlug() ||
			ref == utils.BaseSlug(req.CanonicalSlug()) ||
			ref == utils.Slugify(req.Name) {
			return true
		}
	}
	return false
}

// requirementOverlapFindings uses TF-IDF cosine similarity to find
// semantically similar requirements. Common terms ("syde", "entity",
// "audit", "shall") are naturally down-weighted so only genuinely
// similar statements trigger. Threshold: 0.6 cosine similarity.
func requirementOverlapFindings(all []model.EntityWithBody) []Finding {
	var active []reqEntry
	var allTermSets []map[string]bool
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || req.RequirementStatus != model.RequirementActive || req.Statement == "" {
			continue
		}
		terms := SignificantTerms(req.Statement)
		if len(terms) > 0 {
			active = append(active, reqEntry{req: req, terms: terms})
			allTermSets = append(allTermSets, terms)
		}
	}

	if len(active) < 2 {
		return nil
	}

	corpus := NewTFIDFCorpus(allTermSets)

	audited := func(req *model.RequirementEntity, targetSlug string) bool {
		for _, ao := range req.AuditedOverlaps {
			if ao.Slug == targetSlug || ao.Slug == utils.BaseSlug(targetSlug) {
				return true
			}
		}
		return false
	}

	var out []Finding
	for i := 0; i < len(active); i++ {
		for j := i + 1; j < len(active); j++ {
			a, b := active[i], active[j]
			sim := corpus.TFIDFSimilarity(a.terms, b.terms)
			if sim <= 0.6 {
				continue
			}
			aSlug := a.req.GetBase().CanonicalSlug()
			bSlug := b.req.GetBase().CanonicalSlug()
			pct := sim * 100
			if !audited(a.req, bSlug) {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatRequirement,
					Message:    fmt.Sprintf("overlaps %q (%.0f%% TF-IDF similarity) — add --audited %s to acknowledge", b.req.GetBase().Name, pct, bSlug),
					EntityKind: model.KindRequirement,
					EntitySlug: aSlug,
					EntityName: a.req.GetBase().Name,
					Field:      "audited_overlaps",
				})
			}
			if !audited(b.req, aSlug) {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatRequirement,
					Message:    fmt.Sprintf("overlaps %q (%.0f%% TF-IDF similarity) — add --audited %s to acknowledge", a.req.GetBase().Name, pct, aSlug),
					EntityKind: model.KindRequirement,
					EntitySlug: bSlug,
					EntityName: b.req.GetBase().Name,
					Field:      "audited_overlaps",
				})
			}
		}
	}
	out = append(out, rubberStampOverlapFindings(active)...)
	return out
}

// rubberStampOverlapFindings walks every active requirement's audited
// overlaps and emits an ERROR for each acknowledgement whose distinction
// rationale is empty or trivially short. The audit treats acknowledgement
// without semantic reasoning as unresolved — text similarity is only a
// candidate-surfacing signal, and Claude must author the why.
func rubberStampOverlapFindings(active []reqEntry) []Finding {
	const minDistinction = 20
	var out []Finding
	for _, entry := range active {
		for _, ao := range entry.req.AuditedOverlaps {
			trimmed := strings.TrimSpace(ao.Distinction)
			if len(trimmed) >= minDistinction {
				continue
			}
			reqSlug := entry.req.GetBase().CanonicalSlug()
			msg := fmt.Sprintf("acknowledgement of overlap with %q has no distinction rationale — add --audited %s:\"why this requirement means something different\"", ao.Slug, ao.Slug)
			if trimmed != "" {
				msg = fmt.Sprintf("acknowledgement of overlap with %q has a %d-character distinction rationale; the audit requires at least %d characters of semantic explanation", ao.Slug, len(trimmed), minDistinction)
			}
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatRequirement,
				Message:    msg,
				EntityKind: model.KindRequirement,
				EntitySlug: reqSlug,
				EntityName: entry.req.GetBase().Name,
				Field:      "audited_overlaps",
			})
		}
	}
	return out
}

// requirementContractSurfaceFindings walks every active requirement
// and, for each CLI/REST/screen/event surface named in its statement,
// checks whether some active contract's Input field covers it. When a
// surface has no covering contract, the rule emits an ERROR that names
// the missing surface and suggests authoring a contract or rewording
// the requirement. This is the post-plan mirror of the planning-time
// contractCoverageFindings rule.
func requirementContractSurfaceFindings(all []model.EntityWithBody) []Finding {
	var contractInputs []string
	for _, ewb := range all {
		contract, ok := ewb.Entity.(*model.ContractEntity)
		if !ok {
			continue
		}
		if contract.Input != "" {
			contractInputs = append(contractInputs, contract.Input)
		}
		// Also match on the contract's display name — many flows and
		// requirements reference contracts by their human-readable
		// names (e.g. "Place Order") rather than invocation signatures.
		if contract.Name != "" {
			contractInputs = append(contractInputs, contract.Name)
		}
	}

	var out []Finding
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || req.RequirementStatus != model.RequirementActive || req.Statement == "" {
			continue
		}
		for _, surface := range ExtractSurfaces(req.Statement) {
			covered := false
			for _, input := range contractInputs {
				if ContractCoversSurface(input, surface) {
					covered = true
					break
				}
			}
			if covered {
				continue
			}
			b := req.GetBase()
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatRequirement,
				Message:    fmt.Sprintf("statement names %s surface %q but no active contract's input covers it — author a contract whose input includes this surface, or reword the requirement", surface.Kind, surface.Raw),
				EntityKind: model.KindRequirement,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "statement.surface_coverage",
			})
		}
	}
	return out
}

func requirementFinding(req *model.RequirementEntity, field, msg string) Finding {
	return Finding{
		Severity:   SeverityError,
		Category:   CatRequirement,
		Message:    msg,
		EntityKind: model.KindRequirement,
		EntitySlug: req.CanonicalSlug(),
		EntityName: req.Name,
		Field:      field,
	}
}
