package model

// PhaseAction represents what a plan phase does.
type PhaseAction string

const (
	ActionCreate PhaseAction = "create"
	ActionUpdate PhaseAction = "update"
	ActionDelete PhaseAction = "delete"
)

// PhaseStatus represents the status of a plan phase.
type PhaseStatus string

const (
	PhasePending    PhaseStatus = "pending"
	PhaseInProgress PhaseStatus = "in_progress"
	PhaseCompleted  PhaseStatus = "completed"
	PhaseSkipped    PhaseStatus = "skipped"
)

// PlanPhase is a single phase in a plan. Phases no longer carry draft
// entities — tasks inside the phase instead reference existing entities
// they affect via TaskEntity.AffectedEntities and AffectedFiles.
type PlanPhase struct {
	ID          string      `yaml:"id"`
	Name        string      `yaml:"name,omitempty"`
	ParentPhase string      `yaml:"parent_phase,omitempty"`
	Status      PhaseStatus `yaml:"status"`
	Description string      `yaml:"description,omitempty"`
	Objective   string      `yaml:"objective,omitempty"`
	Changes     string      `yaml:"changes,omitempty"`
	Details     string      `yaml:"details,omitempty"`
	Notes       string      `yaml:"notes,omitempty"`
	Tasks       []string    `yaml:"tasks,omitempty"`
}

// PlanStatus represents the lifecycle of a plan.
type PlanStatus string

const (
	PlanDraft      PlanStatus = "draft"
	PlanApproved   PlanStatus = "approved"
	PlanInProgress PlanStatus = "in-progress"
	PlanCompleted  PlanStatus = "completed"
)

// DeletedChange records an entity a plan will delete.
type DeletedChange struct {
	ID    string   `yaml:"id" json:"id"`
	Slug  string   `yaml:"slug" json:"slug"`
	Why   string   `yaml:"why" json:"why"`
	Tasks []string `yaml:"tasks,omitempty" json:"tasks,omitempty"`
}

// ExtendedChange records an in-place modification to an existing
// entity. FieldChanges is optional; when present it maps frontmatter
// field names to their proposed new value (the special sentinel
// "DELETE" means the field should be cleared). The plan completion
// validator diffs FieldChanges against the target entity's current
// state and errors on any mismatch. Extended entries without
// FieldChanges are hand-review only — the validator emits a WARN
// rather than an ERROR.
type ExtendedChange struct {
	ID           string            `yaml:"id" json:"id"`
	Slug         string            `yaml:"slug" json:"slug"`
	What         string            `yaml:"what" json:"what"`
	Why          string            `yaml:"why" json:"why"`
	FieldChanges map[string]string `yaml:"field_changes,omitempty" json:"field_changes,omitempty"`
	Tasks        []string          `yaml:"tasks,omitempty" json:"tasks,omitempty"`
}

// NewChange drafts a brand-new entity inside a plan. Draft carries
// kind-specific fields (responsibility/capabilities for components,
// input/output/contract_kind/wireframe for contracts, statement/
// req_type/priority/verification for requirements, etc.). The plan
// completion validator expects an entity with the declared Name and
// kind to exist once the plan is marked complete.
type NewChange struct {
	ID    string                 `yaml:"id" json:"id"`
	Name  string                 `yaml:"name" json:"name"`
	What  string                 `yaml:"what" json:"what"`
	Why   string                 `yaml:"why" json:"why"`
	Draft map[string]interface{} `yaml:"draft,omitempty" json:"draft,omitempty"`
	Tasks []string               `yaml:"tasks,omitempty" json:"tasks,omitempty"`
}

// ChangeLane is the per-kind bucket of diff entries for a plan.
type ChangeLane struct {
	Deleted  []DeletedChange  `yaml:"deleted,omitempty" json:"deleted,omitempty"`
	Extended []ExtendedChange `yaml:"extended,omitempty" json:"extended,omitempty"`
	New      []NewChange      `yaml:"new,omitempty" json:"new,omitempty"`
}

// PlanChanges is the structured diff a plan carries: six per-kind
// lanes (requirements/systems/concepts/components/contracts/flows)
// each with their own Deleted/Extended/New lists.
type PlanChanges struct {
	Requirements ChangeLane `yaml:"requirements,omitempty" json:"requirements,omitempty"`
	Systems      ChangeLane `yaml:"systems,omitempty" json:"systems,omitempty"`
	Concepts     ChangeLane `yaml:"concepts,omitempty" json:"concepts,omitempty"`
	Components   ChangeLane `yaml:"components,omitempty" json:"components,omitempty"`
	Contracts    ChangeLane `yaml:"contracts,omitempty" json:"contracts,omitempty"`
	Flows        ChangeLane `yaml:"flows,omitempty" json:"flows,omitempty"`
}

// PlanEntity represents an implementation plan.
type PlanEntity struct {
	BaseEntity     `yaml:",inline"`
	PlanStatus     PlanStatus  `yaml:"plan_status,omitempty"`
	Background     string      `yaml:"background,omitempty"`
	Objective      string      `yaml:"objective,omitempty"`
	PlanScope      string      `yaml:"scope,omitempty"`
	Design         string      `yaml:"design,omitempty"`
	Source         string      `yaml:"source,omitempty"`
	ClaudePlanFile string      `yaml:"claude_plan_file,omitempty"`
	CreatedAt      string      `yaml:"created_at,omitempty"`
	ApprovedAt     string      `yaml:"approved_at,omitempty"`
	CompletedAt    string      `yaml:"completed_at,omitempty"`
	Phases         []PlanPhase `yaml:"phases,omitempty"`
	Changes        PlanChanges `yaml:"changes,omitempty"`
}

// Progress returns the completion percentage of the plan.
func (p *PlanEntity) Progress() float64 {
	if len(p.Phases) == 0 {
		return 0
	}
	completed := 0
	for _, ph := range p.Phases {
		if ph.Status == PhaseCompleted || ph.Status == PhaseSkipped {
			completed++
		}
	}
	return float64(completed) / float64(len(p.Phases)) * 100
}

// PhaseByID returns a phase by its ID, or nil.
func (p *PlanEntity) PhaseByID(id string) *PlanPhase {
	for i := range p.Phases {
		if p.Phases[i].ID == id {
			return &p.Phases[i]
		}
	}
	return nil
}

// ChildPhases returns direct children of a phase. Phases with an
// empty ID are skipped (every phase a caller can recurse into needs
// a real identity), and a phase that lists itself as its own parent
// is skipped too — otherwise CollectTasks would loop forever walking
// "children" that are actually the node it was called with. Defends
// against malformed on-disk data so the model layer never crashes
// the renderer, even if a buggy writer once emitted garbage phases.
func (p *PlanEntity) ChildPhases(parentID string) []PlanPhase {
	var children []PlanPhase
	for _, ph := range p.Phases {
		if ph.ID == "" {
			continue
		}
		if ph.ID == parentID {
			continue
		}
		if ph.ParentPhase == parentID {
			children = append(children, ph)
		}
	}
	return children
}

// CollectTasks returns all tasks for a phase including all descendant
// phases. Uses a visited set so ParentPhase cycles (A->B->A) or
// self-loops cannot cause unbounded recursion — every phase ID is
// walked at most once per call.
func (p *PlanEntity) CollectTasks(phaseID string) []string {
	visited := make(map[string]bool)
	return p.collectTasks(phaseID, visited)
}

// collectTasks is the recursive worker behind CollectTasks. The
// visited map is threaded by reference so sibling branches share
// observations and we never descend into a phase whose ID we have
// already processed on this call.
func (p *PlanEntity) collectTasks(phaseID string, visited map[string]bool) []string {
	if phaseID == "" || visited[phaseID] {
		return nil
	}
	visited[phaseID] = true
	ph := p.PhaseByID(phaseID)
	if ph == nil {
		return nil
	}
	tasks := append([]string{}, ph.Tasks...)
	for _, child := range p.ChildPhases(phaseID) {
		tasks = append(tasks, p.collectTasks(child.ID, visited)...)
	}
	return tasks
}

// AllTasks returns all tasks across all phases.
func (p *PlanEntity) AllTasks() []string {
	var tasks []string
	for _, ph := range p.Phases {
		tasks = append(tasks, ph.Tasks...)
	}
	return tasks
}
