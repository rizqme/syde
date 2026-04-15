package audit

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// requirementFanoutCap is the maximum number of outbound relationships
// a single requirement may receive from entities of the same kind. A
// catch-all "baseline" requirement linking every component in the repo
// satisfies outbound traceability on paper while conveying no intent —
// the cap forces authors to split it into kind-specific requirements.
const requirementFanoutCap = 10

type auditEntity struct {
	Entity model.Entity
	Base   *model.BaseEntity
}

type auditRel struct {
	Source *model.BaseEntity
	Target *model.BaseEntity
	Rel    model.Relationship
}

// requirementTraceFindings flags task/plan/flow entities that have
// no outgoing relationship to a requirement. Kinds already covered
// by coverageFindings (component, contract, concept, system) are
// excluded here — coverage is the stricter check and runs
// separately. Emitted as WARN so deleting the requirement store
// doesn't brick the strict gate for work artefacts.
func requirementTraceFindings(all []model.EntityWithBody) []Finding {
	lookup, _ := auditGraph(all)
	linked := map[string]bool{}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindRequirement {
			continue
		}
		for _, rel := range b.Relationships {
			target, ok := lookup[rel.Target]
			if !ok {
				continue
			}
			if target.Base.Kind == model.KindRequirement {
				linked[b.CanonicalSlug()] = true
				break
			}
		}
	}

	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindRequirement {
			continue
		}
		if coverageStrictKinds[b.Kind] {
			continue // coverage audit handles these
		}
		if linked[b.CanonicalSlug()] {
			continue
		}
		out = append(out, Finding{
			Severity:   SeverityWarning,
			Category:   CatTraceability,
			Message:    "must have an outgoing relationship to at least one requirement",
			EntityKind: b.Kind,
			EntitySlug: b.CanonicalSlug(),
			EntityName: b.Name,
			Field:      "relationships",
		})
	}
	return out
}

// coverageStrictKinds are the entity kinds that MUST be linked to at
// least one requirement (inbound or outbound). Plans, tasks, flows,
// designs are excluded — they're work artefacts or behavior
// descriptions, not properties to verify.
var coverageStrictKinds = map[model.EntityKind]bool{
	model.KindSystem:    true,
	model.KindComponent: true,
	model.KindContract:  true,
	model.KindConcept:   true,
}

// coverageFindings flags any system, component, contract, or concept
// that is not connected to a requirement entity in either direction.
// "Connected" means there exists a relationship edge between the
// entity and any requirement, regardless of relationship type.
func coverageFindings(all []model.EntityWithBody) []Finding {
	lookup, _ := auditGraph(all)

	covered := map[string]bool{}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindRequirement {
			// Outbound from a requirement → mark the target as covered.
			for _, rel := range b.Relationships {
				target, ok := lookup[rel.Target]
				if !ok {
					continue
				}
				if coverageStrictKinds[target.Base.Kind] {
					covered[target.Base.CanonicalSlug()] = true
				}
			}
			continue
		}
		if !coverageStrictKinds[b.Kind] {
			continue
		}
		// Outbound from a strict kind → look for any requirement target.
		for _, rel := range b.Relationships {
			target, ok := lookup[rel.Target]
			if !ok {
				continue
			}
			if target.Base.Kind == model.KindRequirement {
				covered[b.CanonicalSlug()] = true
				break
			}
		}
	}

	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if !coverageStrictKinds[b.Kind] {
			continue
		}
		if covered[b.CanonicalSlug()] {
			continue
		}
		out = append(out, Finding{
			Severity:   SeverityError,
			Category:   CatTraceability,
			Message:    "must be connected to at least one requirement (via refines, derives_from, or any other relationship — in either direction)",
			EntityKind: b.Kind,
			EntitySlug: b.CanonicalSlug(),
			EntityName: b.Name,
			Field:      "relationships",
		})
	}
	return out
}

// goodRequirementFindings checks every requirement entity for the
// fields that make a requirement actually useful: an EARS-compliant
// statement, a req_type, a priority, and a verification description.
// Anything missing or non-conforming surfaces as ERROR — the save
// validator already enforces EARS at write time, but this catches
// hand-edited files and files written before the validator existed.
func goodRequirementFindings(all []model.EntityWithBody) []Finding {
	validTypes := map[model.RequirementType]bool{}
	for _, t := range model.AllRequirementTypes() {
		validTypes[t] = true
	}
	validPriorities := map[model.RequirementPriority]bool{}
	for _, p := range model.AllRequirementPriorities() {
		validPriorities[p] = true
	}

	var out []Finding
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok {
			continue
		}
		b := req.GetBase()
		if _, ok := model.MatchEARS(req.Statement); !ok {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "statement must match an EARS pattern (Ubiquitous, Event-driven, State-driven, Unwanted-behavior, Optional-feature)",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "statement",
			})
		}
		if req.ReqType == "" {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "req_type is required (functional, non-functional, constraint, interface, performance, security, usability)",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "req_type",
			})
		} else if !validTypes[req.ReqType] {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    fmt.Sprintf("req_type %q is not a valid value", req.ReqType),
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "req_type",
			})
		}
		if req.Priority == "" {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "priority is required (must, should, could, wont)",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "priority",
			})
		} else if !validPriorities[req.Priority] {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    fmt.Sprintf("priority %q is not a valid value", req.Priority),
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "priority",
			})
		}
		if strings.TrimSpace(req.Verification) == "" {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "verification is required: describe how the property is verified (test, integration test, inspection, manual)",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "verification",
			})
		}
		if len(b.Files) > 0 {
			out = append(out, Finding{
				Severity:   SeverityError,
				Category:   CatTraceability,
				Message:    "requirements must not list files; link to components/contracts/concepts/systems via refines instead",
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "files",
			})
		}
	}
	return out
}

// requirementFanoutFindings flags requirements that are referenced by
// more than requirementFanoutCap entities of a single source kind.
// Counts outbound non-requirement→requirement relationships, grouped
// by (requirement, source kind). Emits one error per violation.
func requirementFanoutFindings(all []model.EntityWithBody) []Finding {
	lookup, _ := auditGraph(all)

	type key struct {
		reqSlug, reqName string
		sourceKind       model.EntityKind
	}
	counts := map[key]int{}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind == model.KindRequirement {
			continue
		}
		seen := map[string]bool{}
		for _, rel := range b.Relationships {
			target, ok := lookup[rel.Target]
			if !ok {
				continue
			}
			if target.Base.Kind != model.KindRequirement {
				continue
			}
			slug := target.Base.CanonicalSlug()
			if seen[slug] {
				continue
			}
			seen[slug] = true
			counts[key{reqSlug: slug, reqName: target.Base.Name, sourceKind: b.Kind}]++
		}
	}

	var keys []key
	for k := range counts {
		if counts[k] > requirementFanoutCap {
			keys = append(keys, k)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].reqSlug != keys[j].reqSlug {
			return keys[i].reqSlug < keys[j].reqSlug
		}
		return keys[i].sourceKind < keys[j].sourceKind
	})

	var out []Finding
	for _, k := range keys {
		out = append(out, Finding{
			Severity:   SeverityError,
			Category:   CatTraceability,
			Message:    fmt.Sprintf("requirement is linked by %d %s entities (cap is %d per kind) — split into kind-specific requirements", counts[k], k.sourceKind, requirementFanoutCap),
			EntityKind: model.KindRequirement,
			EntitySlug: k.reqSlug,
			EntityName: k.reqName,
			Field:      "relationships",
		})
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
