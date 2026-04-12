package model

import "fmt"

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

	switch v := e.(type) {
	case *ComponentEntity:
		if v.Responsibility == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "responsibility", Message: "recommended"})
		}
	case *ContractEntity:
		if v.ContractKind == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "contract_kind", Message: "recommended"})
		}
	case *FlowEntity:
		if v.Trigger == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "trigger", Message: "recommended"})
		}
	case *DecisionEntity:
		if v.Statement == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "statement", Message: "recommended"})
		}
	case *LearningEntity:
		if v.Category == "" {
			errs = append(errs, ValidationError{EntityID: b.ID, Field: "category", Message: "required"})
		}
	}

	return errs
}
