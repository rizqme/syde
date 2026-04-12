package model

// EntityKind represents the type of design entity.
type EntityKind string

const (
	KindSystem    EntityKind = "system"
	KindComponent EntityKind = "component"
	KindContract  EntityKind = "contract"
	KindConcept   EntityKind = "concept"
	KindFlow      EntityKind = "flow"
	KindDecision  EntityKind = "decision"
	KindPlan      EntityKind = "plan"
	KindTask      EntityKind = "task"
	KindDesign    EntityKind = "design"
	KindLearning  EntityKind = "learning"
)

// AllEntityKinds returns all valid entity kinds.
func AllEntityKinds() []EntityKind {
	return []EntityKind{
		KindSystem, KindComponent, KindContract, KindConcept,
		KindFlow, KindDecision, KindPlan, KindTask, KindDesign, KindLearning,
	}
}

// KindPlural returns the plural directory name for an entity kind.
func (k EntityKind) KindPlural() string {
	switch k {
	case KindSystem:
		return "systems"
	case KindComponent:
		return "components"
	case KindContract:
		return "contracts"
	case KindConcept:
		return "concepts"
	case KindFlow:
		return "flows"
	case KindDecision:
		return "decisions"
	case KindPlan:
		return "plans"
	case KindTask:
		return "tasks"
	case KindDesign:
		return "designs"
	case KindLearning:
		return "learnings"
	default:
		return string(k) + "s"
	}
}

// IDPrefix returns the ID prefix for an entity kind.
func (k EntityKind) IDPrefix() string {
	switch k {
	case KindSystem:
		return "sys"
	case KindComponent:
		return "comp"
	case KindContract:
		return "cont"
	case KindConcept:
		return "conc"
	case KindFlow:
		return "flow"
	case KindDecision:
		return "dec"
	case KindPlan:
		return "plan"
	case KindTask:
		return "task"
	case KindDesign:
		return "des"
	case KindLearning:
		return "learn"
	default:
		return string(k)
	}
}

// ParseEntityKind parses a string into an EntityKind.
func ParseEntityKind(s string) (EntityKind, bool) {
	k := EntityKind(s)
	for _, valid := range AllEntityKinds() {
		if k == valid {
			return k, true
		}
	}
	return "", false
}

// Status represents the lifecycle status of an entity.
type Status string

const (
	StatusDraft      Status = "draft"
	StatusActive     Status = "active"
	StatusDeprecated Status = "deprecated"
	StatusProposed   Status = "proposed"
	StatusSuperseded Status = "superseded"
)

// Relationship represents a typed link between two entities.
type Relationship struct {
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
	Label  string `yaml:"label,omitempty"`
}

// BaseEntity contains fields shared by all entity types.
type BaseEntity struct {
	ID            string         `yaml:"id"`
	Kind          EntityKind     `yaml:"kind"`
	Name          string         `yaml:"name"`
	Description   string         `yaml:"description,omitempty"`
	Purpose       string         `yaml:"purpose,omitempty"`
	Status        Status         `yaml:"status,omitempty"`
	Tags          []string       `yaml:"tags,omitempty"`
	Notes         []string       `yaml:"notes,omitempty"`
	Relationships []Relationship `yaml:"relationships,omitempty"`
}

// GetBase returns the BaseEntity for any entity.
func (b *BaseEntity) GetBase() *BaseEntity { return b }

// Entity is the interface all entity types implement.
type Entity interface {
	GetBase() *BaseEntity
}

// Body holds the markdown body text separate from the frontmatter.
type Body struct {
	Content string
}

// EntityWithBody pairs an entity with its markdown body.
type EntityWithBody struct {
	Entity Entity
	Body   string
}

// SystemEntity represents the top-level system.
type SystemEntity struct {
	BaseEntity       `yaml:",inline"`
	Context          string `yaml:"context,omitempty"`
	Scope            string `yaml:"scope,omitempty"`
	DesignPrinciples string `yaml:"design_principles,omitempty"`
	QualityGoals     string `yaml:"quality_goals,omitempty"`
	Assumptions      string `yaml:"assumptions,omitempty"`
}

// ComponentEntity represents a major architectural part.
type ComponentEntity struct {
	BaseEntity         `yaml:",inline"`
	Responsibility     string   `yaml:"responsibility,omitempty"`
	Boundaries         string   `yaml:"boundaries,omitempty"`
	BehaviorSummary    string   `yaml:"behavior_summary,omitempty"`
	InteractionSummary string   `yaml:"interaction_summary,omitempty"`
	DataHandling       string   `yaml:"data_handling,omitempty"`
	ScalingNotes       string   `yaml:"scaling_notes,omitempty"`
	FailureModes       []string `yaml:"failure_modes,omitempty"`
}

// ContractEntity represents an interaction boundary.
type ContractEntity struct {
	BaseEntity         `yaml:",inline"`
	ContractKind       string `yaml:"contract_kind,omitempty"`
	InteractionPattern string `yaml:"interaction_pattern,omitempty"`
	ProtocolNotes      string `yaml:"protocol_notes,omitempty"`
	InputDescription   string `yaml:"input_description,omitempty"`
	OutputDescription  string `yaml:"output_description,omitempty"`
	Constraints        string `yaml:"constraints,omitempty"`
	VersioningNotes    string `yaml:"versioning_notes,omitempty"`
	UsageExamples      string `yaml:"usage_examples,omitempty"`
}

// ConceptEntity represents a domain/information concept.
type ConceptEntity struct {
	BaseEntity           `yaml:",inline"`
	Meaning              string `yaml:"meaning,omitempty"`
	StructureNotes       string `yaml:"structure_notes,omitempty"`
	Lifecycle            string `yaml:"lifecycle,omitempty"`
	Invariants           string `yaml:"invariants,omitempty"`
	ConceptRelationships string `yaml:"concept_relationships,omitempty"`
	DataSensitivity      string `yaml:"data_sensitivity,omitempty"`
}

// FlowEntity represents meaningful behavior over time.
type FlowEntity struct {
	BaseEntity         `yaml:",inline"`
	Trigger            string `yaml:"trigger,omitempty"`
	Goal               string `yaml:"goal,omitempty"`
	Narrative          string `yaml:"narrative,omitempty"`
	HappyPath          string `yaml:"happy_path,omitempty"`
	EdgeCases          string `yaml:"edge_cases,omitempty"`
	FlowFailureModes   string `yaml:"failure_modes,omitempty"`
	ObservabilityNotes string `yaml:"observability_notes,omitempty"`
	PerformanceNotes   string `yaml:"performance_notes,omitempty"`
}

// DecisionEntity captures architectural intent.
type DecisionEntity struct {
	BaseEntity             `yaml:",inline"`
	Category               string `yaml:"category,omitempty"`
	Statement              string `yaml:"statement,omitempty"`
	Rationale              string `yaml:"rationale,omitempty"`
	AlternativesConsidered string `yaml:"alternatives_considered,omitempty"`
	Tradeoffs              string `yaml:"tradeoffs,omitempty"`
	Consequences           string `yaml:"consequences,omitempty"`
	ReviewNotes            string `yaml:"review_notes,omitempty"`
}

// NewEntityForKind creates a zero-value entity of the given kind.
func NewEntityForKind(kind EntityKind) Entity {
	switch kind {
	case KindSystem:
		return &SystemEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindComponent:
		return &ComponentEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindContract:
		return &ContractEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindConcept:
		return &ConceptEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindFlow:
		return &FlowEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindDecision:
		return &DecisionEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindPlan:
		return &PlanEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindTask:
		return &TaskEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindDesign:
		return &DesignEntity{BaseEntity: BaseEntity{Kind: kind}}
	case KindLearning:
		return &LearningEntity{BaseEntity: BaseEntity{Kind: kind}}
	default:
		return &BaseEntity{Kind: kind}
	}
}
