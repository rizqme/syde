package model

// DesignType represents the type of UI design.
type DesignType string

const (
	DesignScreen    DesignType = "screen"
	DesignFlow      DesignType = "flow"
	DesignComponent DesignType = "component"
)

// DesignEntity represents a UI mockup/design.
type DesignEntity struct {
	BaseEntity    `yaml:",inline"`
	DesignType    DesignType `yaml:"design_type,omitempty"`
	ComponentRefs []string   `yaml:"component_refs,omitempty"`
	FlowRefs      []string   `yaml:"flow_refs,omitempty"`
	ConceptRefs   []string   `yaml:"concept_refs,omitempty"`
}
