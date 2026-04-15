package audit

import (
	"github.com/feedloop/syde/internal/model"
)

// entityFieldFindings runs model.ValidateEntity for each entity and
// converts the per-field errors into audit Findings. "required" becomes
// an Error, "recommended" becomes a Warning.
func entityFieldFindings(all []model.EntityWithBody) []Finding {
	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		for _, ve := range model.ValidateEntity(ewb.Entity) {
			sev := SeverityError
			cat := CatMissingField
			msg := ve.Field + " " + ve.Message
			if ve.Message == "recommended" {
				sev = SeverityWarning
				cat = CatRecommendedField
				msg = ve.Field + " is recommended"
			}
			out = append(out, Finding{
				Severity:   sev,
				Category:   cat,
				Message:    msg,
				EntityKind: b.Kind,
				EntitySlug: b.CanonicalSlug(),
				EntityName: b.Name,
				Field:      ve.Field,
			})
		}
	}
	return out
}
