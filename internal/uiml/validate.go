package uiml

import "fmt"

// ValidationError is a UIML validation issue.
type ValidationError struct {
	Line    int    `json:"line"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.Message)
}

// Validate checks UIML source for errors.
func Validate(source string) []ValidationError {
	result := Parse(source)
	var errs []ValidationError

	// Collect parse errors
	for _, pe := range result.Errors {
		errs = append(errs, ValidationError{Line: pe.Line, Message: pe.Message})
	}

	// Validate nodes
	for _, node := range result.Nodes {
		errs = append(errs, validateNode(node)...)
	}

	return errs
}

func validateNode(node *Node) []ValidationError {
	var errs []ValidationError
	valid := ValidTags()

	// Check tag is known
	if _, ok := valid[string(node.Kind)]; !ok && node.Kind != NodeTextContent {
		errs = append(errs, ValidationError{
			Line:    node.Line,
			Message: fmt.Sprintf("unknown tag: <%s>", node.Kind),
		})
	}

	// Check table has columns
	if node.Kind == NodeTable {
		hasColumns := false
		for _, child := range node.Children {
			if child.Kind == NodeColumns {
				hasColumns = true
				break
			}
		}
		if !hasColumns {
			errs = append(errs, ValidationError{
				Line:    node.Line,
				Message: "<table> should have a <columns> child",
			})
		}
	}

	// Check variant attribute only on button
	if node.HasAttr("variant") && node.Kind != NodeButton {
		errs = append(errs, ValidationError{
			Line:    node.Line,
			Message: fmt.Sprintf("'variant' attribute is only valid on <button>, found on <%s>", node.Kind),
		})
	}

	// Check direction attribute only on layout/navbar
	if node.HasAttr("direction") && node.Kind != NodeLayout && node.Kind != NodeNavbar && node.Kind != NodeTrend {
		errs = append(errs, ValidationError{
			Line:    node.Line,
			Message: fmt.Sprintf("'direction' attribute is only valid on <layout>/<navbar>/<trend>, found on <%s>", node.Kind),
		})
	}

	// Recurse into children
	for _, child := range node.Children {
		errs = append(errs, validateNode(child)...)
	}

	return errs
}
