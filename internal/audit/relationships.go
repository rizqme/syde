package audit

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// relationshipFindings verifies every relationship target resolves to a
// known entity, and that contract belongs_to targets are systems.
func relationshipFindings(all []model.EntityWithBody) []Finding {
	knownTargets := make(map[string]bool)
	targetKind := make(map[string]model.EntityKind)
	add := func(key string, kind model.EntityKind) {
		if key == "" {
			return
		}
		knownTargets[key] = true
		targetKind[key] = kind
	}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		add(b.ID, b.Kind)
		add(b.CanonicalSlug(), b.Kind)
		add(utils.BaseSlug(b.CanonicalSlug()), b.Kind)
		add(utils.Slugify(b.Name), b.Kind)
	}

	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		for _, rel := range b.Relationships {
			if !knownTargets[rel.Target] {
				out = append(out, Finding{
					Severity:   SeverityWarning,
					Category:   CatBrokenRel,
					Message:    fmt.Sprintf("relationship target %q not found", rel.Target),
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
				})
				continue
			}
			if b.Kind == model.KindContract && rel.Type == model.RelBelongsTo {
				if targetKind[rel.Target] != model.KindSystem {
					out = append(out, Finding{
						Severity:   SeverityError,
						Category:   CatContractOwner,
						Message:    fmt.Sprintf("belongs_to target %s is %s, must be system", rel.Target, targetKind[rel.Target]),
						EntityKind: b.Kind,
						EntitySlug: b.CanonicalSlug(),
						EntityName: b.Name,
					})
				}
			}
		}
	}
	return out
}
