package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/feedloop/syde/internal/uiml"
)

// EARSPattern names one of the five Easy Approach to Requirements
// Syntax templates a requirement statement must match.
type EARSPattern string

const (
	EARSUbiquitous       EARSPattern = "ubiquitous"        // The <subject> shall <action>.
	EARSEventDriven      EARSPattern = "event-driven"      // When <trigger>, the <subject> shall <action>.
	EARSStateDriven      EARSPattern = "state-driven"      // While <state>, the <subject> shall <action>.
	EARSUnwantedBehavior EARSPattern = "unwanted-behavior" // If <unwanted condition>, then the <subject> shall <action>.
	EARSOptionalFeature  EARSPattern = "optional-feature"  // Where <feature is included>, the <subject> shall <action>.
)

// earsRegexes anchors each EARS template. Patterns are deliberately
// loose on subject naming and require "shall" so non-shall imperative
// statements ("Add a button.", "Make X visible.") are rejected.
var earsRegexes = []struct {
	pattern EARSPattern
	re      *regexp.Regexp
}{
	{EARSEventDriven, regexp.MustCompile(`(?is)^\s*when\b.+,\s*the\b.+\bshall\b.+\.\s*$`)},
	{EARSStateDriven, regexp.MustCompile(`(?is)^\s*while\b.+,\s*the\b.+\bshall\b.+\.\s*$`)},
	{EARSUnwantedBehavior, regexp.MustCompile(`(?is)^\s*if\b.+,\s*then\s+the\b.+\bshall\b.+\.\s*$`)},
	{EARSOptionalFeature, regexp.MustCompile(`(?is)^\s*where\b.+,\s*the\b.+\bshall\b.+\.\s*$`)},
	// Ubiquitous goes last so the conditional patterns above win
	// when a sentence opens with When/While/If/Where.
	{EARSUbiquitous, regexp.MustCompile(`(?is)^\s*the\b.+\bshall\b.+\.\s*$`)},
}

// MatchEARS returns the matching EARS pattern (and true) for a
// requirement statement, or "" and false if no template matches.
// Statements may end in a period or not; the regexes tolerate
// trailing whitespace.
func MatchEARS(statement string) (EARSPattern, bool) {
	s := strings.TrimSpace(statement)
	if s == "" {
		return "", false
	}
	if !strings.HasSuffix(s, ".") {
		s += "."
	}
	for _, r := range earsRegexes {
		if r.re.MatchString(s) {
			return r.pattern, true
		}
	}
	return "", false
}

// ValidationError represents a single validation issue.
type ValidationError struct {
	EntityID string
	Field    string
	Message  string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s.%s: %s", e.EntityID, e.Field, e.Message)
}

// ValidateEntity checks required fields for an entity.
func ValidateEntity(e Entity) []ValidationError {
	var errs []ValidationError
	b := e.GetBase()

	if b.ID == "" {
		errs = append(errs, ValidationError{EntityID: b.Name, Field: "id", Message: "required"})
	}
	if b.Kind == "" {
		errs = append(errs, ValidationError{EntityID: b.ID, Field: "kind", Message: "required"})
	}
	if b.Name == "" {
		errs = append(errs, ValidationError{EntityID: b.ID, Field: "name", Message: "required"})
	}
	// Description is required on every design-model entity kind so the
	// dashboard list views render meaningful cards. Plans and tasks are
	// workflow artifacts with their own `objective` field — their
	// description requirement is satisfied structurally by objective.
	descriptionRequired := b.Kind != KindPlan && b.Kind != KindTask
	if req, ok := e.(*RequirementEntity); ok && strings.TrimSpace(req.Statement) != "" {
		descriptionRequired = false
	}
	if strings.TrimSpace(b.Description) == "" && descriptionRequired {
		errs = append(errs, ValidationError{EntityID: b.ID, Field: "description", Message: "required"})
	}

	// File references must be concrete paths. No wildcards / globs / braces
	// — each entry is expected to exist as a real node in the tree
	// (.syde/tree.yaml), and wildcard patterns break that link.
	for _, fp := range b.Files {
		if strings.ContainsAny(fp, "*?[]{}") {
			errs = append(errs, ValidationError{
				EntityID: b.ID,
				Field:    "files[" + fp + "]",
				Message:  "wildcard not allowed — list each concrete file path",
			})
		}
	}

	switch v := e.(type) {
	case *ComponentEntity:
		if b.Purpose == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "purpose", Message: "required"})
		}
		if v.Responsibility == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "responsibility", Message: "required"})
		}
		if len(v.Capabilities) == 0 {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "capabilities", Message: "required"})
		}
	case *ContractEntity:
		if v.ContractKind == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "contract_kind", Message: "required"})
		}
		if v.Input == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "input", Message: "required"})
		}
		if len(v.InputParameters) == 0 {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "input_parameters", Message: "required"})
		}
		if v.Output == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "output", Message: "required"})
		}
		if len(v.OutputParameters) == 0 {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "output_parameters", Message: "required"})
		}
		// Contracts must belong_to a system (the process exposing the boundary).
		hasBelongsTo := false
		for _, rel := range b.Relationships {
			if rel.Type == RelBelongsTo {
				hasBelongsTo = true
				break
			}
		}
		if !hasBelongsTo {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "belongs_to", Message: "required"})
		}
		// Screen contracts carry a UIML wireframe source. Require
		// it when contract_kind says so, and run the UIML parser to
		// catch malformed wireframes at save time instead of letting
		// them break the dashboard render later.
		// Screen contracts carry a UIML wireframe. Require it to be
		// non-empty AND parseable through uiml.Parse — every parse
		// error becomes its own ValidationError so the user sees
		// exactly which line is malformed. The UIML lexer is
		// fragile around tag attributes today, so authors should
		// stick to attribute-free structural tags
		// (<screen>/<sidebar>/<main>/<heading>...).
		if v.ContractKind == "screen" {
			if strings.TrimSpace(v.Wireframe) == "" {
				errs = append(errs, ValidationError{EntityID: b.ID, Field: "wireframe", Message: "required — screen contracts must carry a UIML wireframe"})
			} else {
				res := uiml.Parse(v.Wireframe)
				for _, perr := range res.Errors {
					errs = append(errs, ValidationError{EntityID: b.ID, Field: "wireframe", Message: "UIML parse: " + perr.Error()})
				}
			}
		}
	case *FlowEntity:
		if v.Trigger == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "trigger", Message: "recommended"})
		}
	case *RequirementEntity:
		if strings.TrimSpace(v.Statement) == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "statement", Message: "required"})
		} else if _, ok := MatchEARS(v.Statement); !ok {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "statement", Message: "must match an EARS pattern: 'The X shall Y.', 'When Z, the X shall Y.', 'While Z, the X shall Y.', 'If Z, then the X shall Y.', or 'Where Z, the X shall Y.'"})
		}
		switch v.RequirementStatus {
		case RequirementActive:
		case RequirementSuperseded:
			if len(v.SupersededBy) == 0 {
				errs = append(errs, ValidationError{EntityID: b.ID, Field: "superseded_by", Message: "required when requirement_status is superseded"})
			}
		case RequirementObsolete:
			if strings.TrimSpace(v.ObsoleteReason) == "" {
				errs = append(errs, ValidationError{EntityID: b.ID, Field: "obsolete_reason", Message: "required when requirement_status is obsolete"})
			}
		default:
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "requirement_status", Message: "must be active, superseded, or obsolete"})
		}
		switch strings.TrimSpace(v.Source) {
		case "user", "plan", "migration", "manual":
		default:
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "source", Message: "must be user, plan, migration, or manual"})
		}
	case *ConceptEntity:
		if strings.TrimSpace(v.Meaning) == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "meaning", Message: "required"})
		}
		if strings.TrimSpace(v.Invariants) == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "invariants", Message: "recommended"})
		}
	case *PlanEntity:
		errs = append(errs, validatePlanChanges(b.ID, &v.Changes)...)
	}

	return errs
}

// validatePlanChanges walks every lane in a plan's Changes block and
// requires non-empty what/why (or slug/why for Deleted) on every
// entry. NewChange drafts must carry a non-empty kind and at least
// one of the kind-required frontmatter fields; ExtendedChange
// FieldChanges values are free-form strings so only emptiness of
// the map keys is checked. Errors are produced eagerly — a plan
// with several problems reports them all.
func validatePlanChanges(planID string, c *PlanChanges) []ValidationError {
	var errs []ValidationError
	lanes := map[string]ChangeLane{
		"requirements": c.Requirements,
		"systems":      c.Systems,
		"concepts":     c.Concepts,
		"components":   c.Components,
		"contracts":    c.Contracts,
		"flows":        c.Flows,
	}
	for laneName, lane := range lanes {
		for i, d := range lane.Deleted {
			prefix := fmt.Sprintf("changes.%s.deleted[%d]", laneName, i)
			if strings.TrimSpace(d.Slug) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".slug", Message: "required"})
			}
			if strings.TrimSpace(d.Why) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".why", Message: "required"})
			}
		}
		for i, e := range lane.Extended {
			prefix := fmt.Sprintf("changes.%s.extended[%d]", laneName, i)
			if strings.TrimSpace(e.Slug) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".slug", Message: "required"})
			}
			if strings.TrimSpace(e.What) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".what", Message: "required"})
			}
			if strings.TrimSpace(e.Why) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".why", Message: "required"})
			}
			for k := range e.FieldChanges {
				if strings.TrimSpace(k) == "" {
					errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".field_changes", Message: "field key must be non-empty"})
				}
			}
		}
		for i, n := range lane.New {
			prefix := fmt.Sprintf("changes.%s.new[%d]", laneName, i)
			if strings.TrimSpace(n.Name) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".name", Message: "required"})
			}
			if strings.TrimSpace(n.What) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".what", Message: "required"})
			}
			if strings.TrimSpace(n.Why) == "" {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".why", Message: "required"})
			}
			if len(n.Draft) == 0 {
				errs = append(errs, ValidationError{EntityID: planID, Field: prefix + ".draft", Message: "required — NewChange must carry at least one kind-specific field"})
			}
		}
	}
	return errs
}
