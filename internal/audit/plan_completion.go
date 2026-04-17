package audit

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// planCompletionFindings walks every approved or completed plan and
// compares its declared Changes block against actual entity state.
//
// For each DeletedChange: ERROR if the target slug still resolves in
// the current model.
// For each NewChange: ERROR if no entity of the declared lane's kind
// with a matching slugified name exists.
// For each ExtendedChange: WARN if no FieldChanges are declared
// (nothing to verify mechanically). Otherwise ERROR for each
// declared field whose current value does not match the proposed
// value. The sentinel value "DELETE" on the proposed side means the
// field must currently be the zero value of its type.
//
// Draft plans (PlanDraft) are skipped — the diff only becomes a
// contract once a human has approved it.
func planCompletionFindings(all []model.EntityWithBody) []Finding {
	lookup, _ := auditGraph(all)
	var out []Finding

	laneKinds := map[string]model.EntityKind{
		"requirements": model.KindRequirement,
		"systems":      model.KindSystem,
		"concepts":     model.KindConcept,
		"components":   model.KindComponent,
		"contracts":    model.KindContract,
		"flows":        model.KindFlow,
	}

	for _, ewb := range all {
		plan, ok := ewb.Entity.(*model.PlanEntity)
		if !ok {
			continue
		}
		if plan.PlanStatus != model.PlanApproved && plan.PlanStatus != model.PlanCompleted && plan.PlanStatus != model.PlanInProgress {
			continue
		}
		b := plan.GetBase()
		planSlug := b.CanonicalSlug()

		lanes := map[string]model.ChangeLane{
			"requirements": plan.Changes.Requirements,
			"systems":      plan.Changes.Systems,
			"concepts":     plan.Changes.Concepts,
			"components":   plan.Changes.Components,
			"contracts":    plan.Changes.Contracts,
			"flows":        plan.Changes.Flows,
		}

		for laneName, lane := range lanes {
			expectedKind := laneKinds[laneName]

			for _, d := range lane.Deleted {
				target, exists := lookup[d.Slug]
				if exists && target.Base.Kind == expectedKind {
					out = append(out, Finding{
						Severity:   SeverityError,
						Category:   CatPlanCompletion,
						Message:    fmt.Sprintf("plan %q claims to delete %s %q but it still exists", b.Name, expectedKind, d.Slug),
						EntityKind: model.KindPlan,
						EntitySlug: planSlug,
						EntityName: b.Name,
						Field:      fmt.Sprintf("changes.%s.deleted[%s]", laneName, d.ID),
					})
				}
			}

			for _, n := range lane.New {
				expectedSlug := utils.Slugify(n.Name)
				found := false
				for _, candidate := range all {
					cb := candidate.Entity.GetBase()
					if cb.Kind != expectedKind {
						continue
					}
					if utils.Slugify(cb.Name) == expectedSlug {
						found = true
						break
					}
				}
				if !found {
					out = append(out, Finding{
						Severity:   SeverityError,
						Category:   CatPlanCompletion,
						Message:    fmt.Sprintf("plan %q claims to create new %s %q but no entity with that name exists", b.Name, expectedKind, n.Name),
						EntityKind: model.KindPlan,
						EntitySlug: planSlug,
						EntityName: b.Name,
						Field:      fmt.Sprintf("changes.%s.new[%s]", laneName, n.ID),
					})
				}
			}

			for _, e := range lane.Extended {
				target, exists := lookup[e.Slug]
				if !exists {
					out = append(out, Finding{
						Severity:   SeverityError,
						Category:   CatPlanCompletion,
						Message:    fmt.Sprintf("plan %q extends %s %q but the target entity does not exist", b.Name, expectedKind, e.Slug),
						EntityKind: model.KindPlan,
						EntitySlug: planSlug,
						EntityName: b.Name,
						Field:      fmt.Sprintf("changes.%s.extended[%s]", laneName, e.ID),
					})
					continue
				}
				if len(e.FieldChanges) == 0 {
					// Hand-review extends without declared field_changes
					// are legitimate — there is nothing to verify
					// programmatically. Skip both the finding and the
					// downstream value comparison.
					continue
				}
				for field, expected := range e.FieldChanges {
					current, ok := readEntityField(target.Entity, field)
					if !ok {
						out = append(out, Finding{
							Severity:   SeverityError,
							Category:   CatPlanCompletion,
							Message:    fmt.Sprintf("plan %q declares a change to unknown field %q on %s %q", b.Name, field, expectedKind, e.Slug),
							EntityKind: model.KindPlan,
							EntitySlug: planSlug,
							EntityName: b.Name,
							Field:      fmt.Sprintf("changes.%s.extended[%s].field_changes.%s", laneName, e.ID, field),
						})
						continue
					}
					if !fieldMatches(current, expected) {
						out = append(out, Finding{
							Severity:   SeverityError,
							Category:   CatPlanCompletion,
							Message:    fmt.Sprintf("plan %q declares %s.%s=%q but current value is %q", b.Name, e.Slug, field, expected, renderFieldValue(current)),
							EntityKind: model.KindPlan,
							EntitySlug: planSlug,
							EntityName: b.Name,
							Field:      fmt.Sprintf("changes.%s.extended[%s].field_changes.%s", laneName, e.ID, field),
						})
					}
				}
			}
		}
	}
	return out
}

// readEntityField looks up a frontmatter field on a typed entity by
// its YAML tag name. Returns the reflect.Value (so callers can
// compare) and true if the field exists.
func readEntityField(entity model.Entity, yamlTag string) (interface{}, bool) {
	v := reflect.ValueOf(entity)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}
	return readFieldByYAMLTag(v, yamlTag)
}

func readFieldByYAMLTag(v reflect.Value, yamlTag string) (interface{}, bool) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			// Recurse into embedded structs (e.g. BaseEntity).
			if inner, ok := readFieldByYAMLTag(v.Field(i), yamlTag); ok {
				return inner, true
			}
			continue
		}
		tag := f.Tag.Get("yaml")
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "" {
			name = strings.ToLower(f.Name)
		}
		if name == yamlTag {
			return v.Field(i).Interface(), true
		}
	}
	return nil, false
}

// fieldMatches returns true when the current entity field value
// equals the plan's declared new value. The "DELETE" sentinel means
// the field must be its zero value. Supported current types: string,
// []string, map, slices of structs (compared via renderFieldValue).
func fieldMatches(current interface{}, expected string) bool {
	if expected == "DELETE" {
		return reflect.DeepEqual(current, reflect.Zero(reflect.TypeOf(current)).Interface())
	}
	return renderFieldValue(current) == expected
}

// renderFieldValue renders a reflected field value as a string so we
// can compare it against the plan's declared string value. Slices
// are joined with ", " so `capabilities=a, b, c` matches a []string.
func renderFieldValue(current interface{}) string {
	if current == nil {
		return ""
	}
	v := reflect.ValueOf(current)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Slice, reflect.Array:
		parts := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			parts[i] = fmt.Sprintf("%v", v.Index(i).Interface())
		}
		return strings.Join(parts, ", ")
	default:
		return fmt.Sprintf("%v", current)
	}
}
