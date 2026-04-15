package model

import "strings"

// EntityKind represents the type of design entity.
type EntityKind string

const (
	KindSystem      EntityKind = "system"
	KindComponent   EntityKind = "component"
	KindContract    EntityKind = "contract"
	KindConcept     EntityKind = "concept"
	KindFlow        EntityKind = "flow"
	KindDecision    EntityKind = "decision"
	KindPlan        EntityKind = "plan"
	KindTask        EntityKind = "task"
	KindDesign      EntityKind = "design"
	KindLearning    EntityKind = "learning"
	KindRequirement EntityKind = "requirement"
)

// AllEntityKinds returns all valid entity kinds.
func AllEntityKinds() []EntityKind {
	return []EntityKind{
		KindSystem, KindComponent, KindContract, KindConcept,
		KindFlow, KindDecision, KindPlan, KindTask, KindDesign, KindLearning,
		KindRequirement,
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
	case KindRequirement:
		return "requirements"
	default:
		return string(k) + "s"
	}
}

// IDPrefix returns the short uppercase prefix used in entity IDs
// (SYS-0001, COM-0002, etc.). Prefixes are exactly three characters.
func (k EntityKind) IDPrefix() string {
	switch k {
	case KindSystem:
		return "SYS"
	case KindComponent:
		return "COM"
	case KindContract:
		return "CON"
	case KindConcept:
		return "CPT"
	case KindFlow:
		return "FLW"
	case KindDecision:
		return "DEC"
	case KindPlan:
		return "PLN"
	case KindTask:
		return "TSK"
	case KindDesign:
		return "DSG"
	case KindLearning:
		return "LRN"
	case KindRequirement:
		return "REQ"
	default:
		return strings.ToUpper(string(k))
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

// Relationship represents a typed link between two entities.
type Relationship struct {
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
	Label  string `yaml:"label,omitempty"`
}

// BaseEntity contains fields shared by all entity types.
type BaseEntity struct {
	ID   string     `yaml:"id"`
	Kind EntityKind `yaml:"kind"`
	Name string     `yaml:"name"`
	// Slug is the file-level identifier: <name-slugified>-<rand4>,
	// e.g. "cli-a3f2". Generated on create and never changes so files
	// stay stable across renames. Callers can address an entity by the
	// full Slug (always unique), by the bare name-slugified prefix
	// (works when unique), or by "<parent-slug>/<child-slug>" (walks
	// belongs_to to disambiguate).
	Slug             string         `yaml:"slug,omitempty"`
	Description      string         `yaml:"description,omitempty"`
	Purpose          string         `yaml:"purpose,omitempty"`
	Deprecated       bool           `yaml:"deprecated,omitempty"`
	DeprecatedReason string         `yaml:"deprecated_reason,omitempty"`
	ReplacedBy       string         `yaml:"replaced_by,omitempty"`
	Tags             []string       `yaml:"tags,omitempty"`
	Notes            []string       `yaml:"notes,omitempty"`
	Files            []string       `yaml:"files,omitempty"`
	Relationships    []Relationship `yaml:"relationships,omitempty"`
	// UpdatedAt records the last time a task or explicit update touched
	// this entity. Used by the drift validator: if a file referenced in
	// entity.Files was modified (tree node's Mtime) after this entity's
	// UpdatedAt, the entity is flagged as potentially out-of-date.
	UpdatedAt string `yaml:"updated_at,omitempty"`
}

// GetBase returns the BaseEntity for any entity.
func (b *BaseEntity) GetBase() *BaseEntity { return b }

// CanonicalSlug returns the file-addressable slug. Prefers the stored
// b.Slug (suffixed form) but falls back to a simple slugify of the name
// for legacy entities without a Slug field yet.
func (b *BaseEntity) CanonicalSlug() string {
	if b.Slug != "" {
		return b.Slug
	}
	return simpleSlug(b.Name)
}

// simpleSlug is an internal mirror of utils.Slugify to avoid the model
// package importing utils. Lowercases, replaces whitespace with dashes,
// and strips anything non-[a-z0-9-].
func simpleSlug(name string) string {
	var b []byte
	prevDash := false
	for _, r := range name {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b = append(b, byte(r))
			prevDash = false
			continue
		}
		if !prevDash && len(b) > 0 {
			b = append(b, '-')
			prevDash = true
		}
	}
	for len(b) > 0 && b[len(b)-1] == '-' {
		b = b[:len(b)-1]
	}
	return string(b)
}

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
	Capabilities       []string `yaml:"capabilities,omitempty"`
	Boundaries         string   `yaml:"boundaries,omitempty"`
	BehaviorSummary    string   `yaml:"behavior_summary,omitempty"`
	InteractionSummary string   `yaml:"interaction_summary,omitempty"`
	DataHandling       string   `yaml:"data_handling,omitempty"`
	ScalingNotes       string   `yaml:"scaling_notes,omitempty"`
	FailureModes       []string `yaml:"failure_modes,omitempty"`
}

// ContractParam is a single structured input or output parameter. Path is a
// JSON-path-style locator (e.g. "file.path", "items[].name", or just a bare
// name for top-level params). Type is a short type hint (e.g. "string", "int",
// "array<User>"). Description is a short human explanation.
type ContractParam struct {
	Path        string `yaml:"path" json:"path"`
	Type        string `yaml:"type,omitempty" json:"type,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// ParseContractParam parses a "path|type|description" spec into a ContractParam.
// Description may contain pipes. Type is optional. Path is required.
func ParseContractParam(spec string) (ContractParam, bool) {
	parts := strings.SplitN(spec, "|", 3)
	if parts[0] == "" {
		return ContractParam{}, false
	}
	p := ContractParam{Path: strings.TrimSpace(parts[0])}
	if len(parts) > 1 {
		p.Type = strings.TrimSpace(parts[1])
	}
	if len(parts) > 2 {
		p.Description = strings.TrimSpace(parts[2])
	}
	return p, true
}

// ContractEntity represents an interaction boundary.
type ContractEntity struct {
	BaseEntity         `yaml:",inline"`
	ContractKind       string          `yaml:"contract_kind,omitempty"`
	InteractionPattern string          `yaml:"interaction_pattern,omitempty"`
	ProtocolNotes      string          `yaml:"protocol_notes,omitempty"`
	Input              string          `yaml:"input,omitempty"`             // invocation signature (e.g. "GET /api/projects", "syde plan create <name>")
	InputParameters    []ContractParam `yaml:"input_parameters,omitempty"`  // structured input parameters
	Output             string          `yaml:"output,omitempty"`            // output signature / response shape (e.g. "200 OK application/json")
	OutputParameters   []ContractParam `yaml:"output_parameters,omitempty"` // structured output parameters
	Constraints        string          `yaml:"constraints,omitempty"`
	VersioningNotes    string          `yaml:"versioning_notes,omitempty"`
	UsageExamples      string          `yaml:"usage_examples,omitempty"`
	// Wireframe holds UIML source describing the screen layout for
	// screen-kind contracts (contract_kind == "screen"). Validator
	// runs uiml.Parse on it at save time; renderers use
	// uiml.RenderHTML to show a wireframe preview in the dashboard
	// detail panel.
	Wireframe string `yaml:"wireframe,omitempty"`
}

// ConceptAttribute is a single high-level ERD attribute on a concept
// entity. Concepts live at the design level — concrete Go / SQL /
// TypeScript types belong in code, not in the model — so the
// attribute carries only a name, a prose description, and an
// optional list of concept slugs this attribute references.
// Name is required (the field / column / property name), Description
// is free-form prose for invariants / defaults / units / notes, and
// Refs is used by the ERD view to draw attribute-level FK-style
// edges from the attribute row to the referenced concept's card.
type ConceptAttribute struct {
	Name        string   `yaml:"name" json:"name"`
	Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	Refs        []string `yaml:"refs,omitempty" json:"refs,omitempty"`
}

// ConceptAction is a design-level verb on a concept — a domain
// operation like Order.cancel() or User.verifyEmail(). Intentionally
// simpler than a Contract (no input/output parameter schemas):
// actions describe behaviour at the aggregate level, and the full
// interaction surface is modelled separately via contract entities
// when needed.
type ConceptAction struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// ParseConceptAttribute parses a pipe-separated attribute spec:
//
//	name[|description[|ref1,ref2,...]]
//
// Name is required; description and refs are optional. Refs is a
// comma-separated list of concept slugs this attribute points at
// (FK-style). Concepts are a design-level lens so there is no type
// hint — concrete types live in the code.
func ParseConceptAttribute(spec string) (ConceptAttribute, bool) {
	parts := strings.SplitN(spec, "|", 3)
	if strings.TrimSpace(parts[0]) == "" {
		return ConceptAttribute{}, false
	}
	a := ConceptAttribute{Name: strings.TrimSpace(parts[0])}
	if len(parts) > 1 {
		a.Description = strings.TrimSpace(parts[1])
	}
	if len(parts) > 2 {
		for _, r := range strings.Split(parts[2], ",") {
			r = strings.TrimSpace(r)
			if r != "" {
				a.Refs = append(a.Refs, r)
			}
		}
	}
	return a, true
}

// ParseConceptAction parses a "name|description" spec.
func ParseConceptAction(spec string) (ConceptAction, bool) {
	parts := strings.SplitN(spec, "|", 2)
	if strings.TrimSpace(parts[0]) == "" {
		return ConceptAction{}, false
	}
	a := ConceptAction{Name: strings.TrimSpace(parts[0])}
	if len(parts) > 1 {
		a.Description = strings.TrimSpace(parts[1])
	}
	return a, true
}

// ConceptEntity represents a domain/information concept. Now ERD-
// shaped: Attributes list the typed fields, Actions list the domain
// verbs, and relates_to relationships carry a cardinality label
// (one-to-one / one-to-many / many-to-one / many-to-many) via the
// shared Relationship.Label field.
type ConceptEntity struct {
	BaseEntity           `yaml:",inline"`
	Meaning              string             `yaml:"meaning,omitempty"`
	StructureNotes       string             `yaml:"structure_notes,omitempty"`
	Lifecycle            string             `yaml:"lifecycle,omitempty"`
	Invariants           string             `yaml:"invariants,omitempty"`
	ConceptRelationships string             `yaml:"concept_relationships,omitempty"`
	DataSensitivity      string             `yaml:"data_sensitivity,omitempty"`
	Attributes           []ConceptAttribute `yaml:"attributes,omitempty" json:"attributes,omitempty"`
	Actions              []ConceptAction    `yaml:"actions,omitempty" json:"actions,omitempty"`
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
	Supersedes             string `yaml:"supersedes,omitempty"`
}

// RequirementStatus describes whether a requirement is currently
// active or retained only as historical context.
type RequirementStatus string

const (
	RequirementActive     RequirementStatus = "active"
	RequirementSuperseded RequirementStatus = "superseded"
	RequirementObsolete   RequirementStatus = "obsolete"
)

// RequirementEntity captures a user- or plan-approved requirement as
// immutable design intent. Requirements are append-only audit records:
// conflicts are represented by supersedes / superseded_by links and
// status changes, not by deleting historical requirement files.
type RequirementEntity struct {
	BaseEntity         `yaml:",inline"`
	Statement          string            `yaml:"statement,omitempty"`
	Source             string            `yaml:"source,omitempty"`     // user, plan, migration, manual
	SourceRef          string            `yaml:"source_ref,omitempty"` // plan slug, transcript ref, issue URL, etc.
	RequirementStatus  RequirementStatus `yaml:"requirement_status,omitempty"`
	Rationale          string            `yaml:"rationale,omitempty"`
	AcceptanceCriteria string            `yaml:"acceptance_criteria,omitempty"`
	Supersedes         []string          `yaml:"supersedes,omitempty"`
	SupersededBy       []string          `yaml:"superseded_by,omitempty"`
	ObsoleteReason     string            `yaml:"obsolete_reason,omitempty"`
	ApprovedAt         string            `yaml:"approved_at,omitempty"`
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
	case KindRequirement:
		return &RequirementEntity{
			BaseEntity:        BaseEntity{Kind: kind},
			RequirementStatus: RequirementActive,
		}
	default:
		return &BaseEntity{Kind: kind}
	}
}
