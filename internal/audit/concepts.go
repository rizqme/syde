package audit

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
)

// validCardinalities is the closed set of cardinality labels allowed
// on relates_to relationships between concepts. Empty label is
// always permitted — cardinality is optional, not every relates_to
// link needs one. Anything else is rejected so the dashboard ERD
// view can render every label with confidence.
var validCardinalities = map[string]bool{
	"one-to-one":   true,
	"one-to-many":  true,
	"many-to-one":  true,
	"many-to-many": true,
}

// conceptFindings validates ERD-level structure on concept entities:
// every relates_to relationship with a non-empty Label must use one
// of the four canonical cardinality values. Anything else is an
// ERROR so `syde sync check --strict` blocks session end until the
// model is repaired. The required-field checks (meaning, attributes,
// attribute name/type) live in model.ValidateEntity and flow through
// the entityFieldFindings path — this file only covers the cross-
// relationship checks the per-entity validator cannot see.
func conceptFindings(all []model.EntityWithBody) []Finding {
	var out []Finding
	for _, ewb := range all {
		c, ok := ewb.Entity.(*model.ConceptEntity)
		if !ok {
			continue
		}
		b := c.GetBase()
		for _, rel := range b.Relationships {
			if rel.Type != model.RelRelatesTo {
				continue
			}
			label := strings.TrimSpace(rel.Label)
			if label == "" {
				continue
			}
			if !validCardinalities[label] {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatConceptIntegrity,
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Message: fmt.Sprintf(
						"concept %q relates_to %q has invalid cardinality label %q — expected one of one-to-one, one-to-many, many-to-one, many-to-many",
						b.Name, rel.Target, label,
					),
				})
			}
		}
	}
	return out
}
