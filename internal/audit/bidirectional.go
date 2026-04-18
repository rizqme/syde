package audit

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// bidirectionalFindings enforces the requirement↔component coupling
// introduced by plan PLN-0018: every active requirement must refine
// ≥1 component, every component with files mapped must have ≥1
// refining active requirement, and no active requirement may target
// a system via refines or belongs_to. Stale-hash drift detection
// lives in requirementStaleFindings below.
func bidirectionalFindings(all []model.EntityWithBody) []Finding {
	components := map[string]*model.ComponentEntity{}
	systems := map[string]bool{}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		switch e := ewb.Entity.(type) {
		case *model.ComponentEntity:
			components[b.ID] = e
			components[b.CanonicalSlug()] = e
			components[utils.BaseSlug(b.CanonicalSlug())] = e
			components[utils.Slugify(b.Name)] = e
		case *model.SystemEntity:
			_ = e
			systems[b.ID] = true
			systems[b.CanonicalSlug()] = true
			systems[utils.BaseSlug(b.CanonicalSlug())] = true
			systems[utils.Slugify(b.Name)] = true
		}
	}

	var out []Finding

	// requirement_no_component + requirement_targets_system
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || req.RequirementStatus != model.RequirementActive {
			continue
		}
		b := req.GetBase()
		refinesComponent := false
		for _, rel := range b.Relationships {
			switch rel.Type {
			case model.RelRefines:
				if _, ok := components[rel.Target]; ok {
					refinesComponent = true
				}
				if systems[rel.Target] {
					out = append(out, Finding{
						Severity:   SeverityFinding,
						Category:   CatRequirement,
						Message:    fmt.Sprintf("requirement targets system %q via refines — requirements must refine components, not systems", rel.Target),
						EntityKind: model.KindRequirement,
						EntitySlug: b.CanonicalSlug(),
						EntityName: b.Name,
						Field:      "relationships.refines",
					})
				}
			case model.RelBelongsTo:
				if systems[rel.Target] {
					out = append(out, Finding{
						Severity:   SeverityFinding,
						Category:   CatRequirement,
						Message:    fmt.Sprintf("requirement belongs_to system %q — requirements must refine a component, drop the belongs_to:system edge", rel.Target),
						EntityKind: model.KindRequirement,
						EntitySlug: b.CanonicalSlug(),
						EntityName: b.Name,
						Field:      "relationships.belongs_to",
					})
				}
			}
		}
		if !refinesComponent {
			out = append(out, Finding{
				Severity:   SeverityFinding,
				Category:   CatRequirement,
				Message:    "active requirement has no refines:component edge — every requirement must refine at least one component (see plan PLN-0018)",
				EntityKind: model.KindRequirement,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "relationships.refines",
			})
		}
	}

	// component_no_requirement: walk components; for each with files,
	// scan all active reqs for a refines edge landing on any of its
	// aliases. Emit finding when none exist.
	seenComp := map[string]bool{}
	for _, ewb := range all {
		comp, ok := ewb.Entity.(*model.ComponentEntity)
		if !ok {
			continue
		}
		b := comp.GetBase()
		if seenComp[b.ID] {
			continue
		}
		seenComp[b.ID] = true
		if len(b.Files) == 0 {
			continue
		}
		aliases := map[string]bool{
			b.ID:                                true,
			b.CanonicalSlug():                   true,
			utils.BaseSlug(b.CanonicalSlug()):   true,
			utils.Slugify(b.Name):               true,
		}
		found := false
		for _, other := range all {
			req, ok := other.Entity.(*model.RequirementEntity)
			if !ok || req.RequirementStatus != model.RequirementActive {
				continue
			}
			for _, rel := range req.GetBase().Relationships {
				if rel.Type == model.RelRefines && aliases[rel.Target] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			out = append(out, Finding{
				Severity:   SeverityFinding,
				Category:   CatTraceability,
				Message:    fmt.Sprintf("component has %d file(s) mapped but no active requirement refines it — author at least one refining requirement", len(b.Files)),
				EntityKind: model.KindComponent,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      "relationships.refines_incoming",
			})
		}
	}
	return out
}

// requirementStaleFindings emits a requirement_stale finding for every
// refining-component of every active requirement whose current file
// content hash differs from the requirement's stored VerifiedAgainst
// snapshot (or no snapshot exists at all). projectRoot is the absolute
// path of the syde project so relative component file paths can be
// resolved.
func requirementStaleFindings(all []model.EntityWithBody, projectRoot string) []Finding {
	if projectRoot == "" {
		return nil
	}
	components := map[string]*model.ComponentEntity{}
	for _, ewb := range all {
		comp, ok := ewb.Entity.(*model.ComponentEntity)
		if !ok {
			continue
		}
		b := comp.GetBase()
		components[b.ID] = comp
		components[b.CanonicalSlug()] = comp
		components[utils.BaseSlug(b.CanonicalSlug())] = comp
		components[utils.Slugify(b.Name)] = comp
	}

	// Memoise combined-hash-per-component so the same component shared
	// by many requirements is only hashed once per audit run.
	hashCache := map[string]string{}
	hashFor := func(comp *model.ComponentEntity) (string, error) {
		slug := comp.GetBase().CanonicalSlug()
		if h, ok := hashCache[slug]; ok {
			return h, nil
		}
		paths := make([]string, len(comp.GetBase().Files))
		for i, f := range comp.GetBase().Files {
			paths[i] = filepath.Join(projectRoot, f)
		}
		// Sort so insertion order can't shift the hash between runs.
		sort.Strings(paths)
		h, err := utils.CombinedFilesSHA256(paths)
		if err != nil {
			return "", err
		}
		hashCache[slug] = h
		return h, nil
	}

	var out []Finding
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || req.RequirementStatus != model.RequirementActive {
			continue
		}
		b := req.GetBase()
		for _, rel := range b.Relationships {
			if rel.Type != model.RelRefines {
				continue
			}
			comp, ok := components[rel.Target]
			if !ok {
				continue
			}
			if len(comp.GetBase().Files) == 0 {
				// Design-phase component (no files yet) — snapshot is
				// meaningless, so skip stale-check for this edge.
				continue
			}
			compSlug := comp.GetBase().CanonicalSlug()
			current, err := hashFor(comp)
			if err != nil {
				out = append(out, Finding{
					Severity:   SeverityFinding,
					Category:   CatRequirement,
					Message:    fmt.Sprintf("cannot hash files of refining component %q: %v", compSlug, err),
					EntityKind: model.KindRequirement,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Field:      "verified_against",
				})
				continue
			}
			snap, has := req.VerifiedAgainst[compSlug]
			if !has || snap.Hash == "" {
				out = append(out, Finding{
					Severity:   SeverityFinding,
					Category:   CatRequirement,
					Message:    fmt.Sprintf("requirement has no verified_against snapshot for refining component %q — run 'syde requirement verify %s' after reading the requirement", compSlug, b.CanonicalSlug()),
					EntityKind: model.KindRequirement,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Field:      "verified_against",
				})
				continue
			}
			if snap.Hash != current {
				out = append(out, Finding{
					Severity:   SeverityFinding,
					Category:   CatRequirement,
					Message:    fmt.Sprintf("refining component %q has drifted from last verified snapshot (at %s) — re-read the requirement then run 'syde requirement verify %s'", compSlug, snap.At, b.CanonicalSlug()),
					EntityKind: model.KindRequirement,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Field:      "verified_against",
				})
			}
		}
	}
	return out
}
