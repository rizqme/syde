package audit

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

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
