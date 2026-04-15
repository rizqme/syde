package audit

import (
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

type auditEntity struct {
	Entity model.Entity
	Base   *model.BaseEntity
}

type auditRel struct {
	Source *model.BaseEntity
	Target *model.BaseEntity
	Rel    model.Relationship
}

func requirementTraceFindings(all []model.EntityWithBody) []Finding {
	_, rels := auditGraph(all)
	linked := map[string]bool{}
	for _, rel := range rels {
		if rel.Source.Kind == model.KindRequirement && rel.Target.Kind != model.KindRequirement {
			linked[rel.Target.CanonicalSlug()] = true
		}
		if rel.Target.Kind == model.KindRequirement && rel.Source.Kind != model.KindRequirement {
			linked[rel.Source.CanonicalSlug()] = true
		}
	}

	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindRequirement {
			continue
		}
		if !linked[b.CanonicalSlug()] {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "must link to at least one requirement",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "relationships",
			})
		}
	}
	return out
}

func hierarchyFindings(all []model.EntityWithBody) []Finding {
	var roots []*model.BaseEntity
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindSystem && !hasAuditRelType(b.Relationships, model.RelBelongsTo) {
			roots = append(roots, b)
		}
	}

	var out []Finding
	if len(roots) == 0 {
		out = append(out, Finding{
			Severity: SeverityError,
			Category: CatHierarchy,
			Message:  "one root system must exist without belongs_to",
			Field:    "belongs_to",
		})
	}
	if len(roots) > 1 {
		for _, root := range roots {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatHierarchy,
				Message:    "only one root system may omit belongs_to",
				EntityKind: root.Kind,
				EntitySlug: root.CanonicalSlug(),
				EntityName: root.Name,
				Field:      "belongs_to",
			})
		}
	}

	rootSlug := ""
	if len(roots) == 1 {
		rootSlug = roots[0].CanonicalSlug()
	}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindSystem && b.CanonicalSlug() == rootSlug {
			continue
		}
		if !hasAuditRelType(b.Relationships, model.RelBelongsTo) {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatHierarchy,
				Message:    "must have a belongs_to parent",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "belongs_to",
			})
		}
	}
	return out
}

func contractFlowFindings(all []model.EntityWithBody) []Finding {
	_, rels := auditGraph(all)
	inFlow := map[string]bool{}
	for _, rel := range rels {
		if rel.Source.Kind == model.KindContract && rel.Target.Kind == model.KindFlow {
			inFlow[rel.Source.CanonicalSlug()] = true
		}
		if rel.Source.Kind == model.KindFlow && rel.Target.Kind == model.KindContract {
			inFlow[rel.Target.CanonicalSlug()] = true
		}
	}

	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind != model.KindContract {
			continue
		}
		if !inFlow[b.CanonicalSlug()] {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatContractFlow,
				Message:    "contract must participate in at least one flow",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "relationships",
			})
		}
	}
	return out
}

func auditGraph(all []model.EntityWithBody) (map[string]auditEntity, []auditRel) {
	lookup := make(map[string]auditEntity)
	add := func(key string, ent auditEntity) {
		if key != "" {
			lookup[key] = ent
		}
	}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		ent := auditEntity{Entity: ewb.Entity, Base: b}
		add(b.ID, ent)
		add(b.CanonicalSlug(), ent)
		add(utils.BaseSlug(b.CanonicalSlug()), ent)
		add(utils.Slugify(b.Name), ent)
	}

	var rels []auditRel
	for _, ewb := range all {
		source := ewb.Entity.GetBase()
		for _, rel := range source.Relationships {
			target, ok := lookup[rel.Target]
			if !ok {
				continue
			}
			rels = append(rels, auditRel{Source: source, Target: target.Base, Rel: rel})
		}
	}
	return lookup, rels
}

func hasAuditRelType(rels []model.Relationship, relType string) bool {
	for _, rel := range rels {
		if rel.Type == relType {
			return true
		}
	}
	return false
}
