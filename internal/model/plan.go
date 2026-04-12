package model

// PlanStepAction represents what a plan step does.
type PlanStepAction string

const (
	ActionCreate PlanStepAction = "create"
	ActionUpdate PlanStepAction = "update"
	ActionDelete PlanStepAction = "delete"
)

// PlanStepStatus represents the status of a plan step.
type PlanStepStatus string

const (
	StepPending    PlanStepStatus = "pending"
	StepInProgress PlanStepStatus = "in_progress"
	StepCompleted  PlanStepStatus = "completed"
	StepSkipped    PlanStepStatus = "skipped"
)

// PlanStep is a single step in a plan.
type PlanStep struct {
	ID          string         `yaml:"id"`
	Action      PlanStepAction `yaml:"action"`
	EntityKind  EntityKind     `yaml:"entity_kind"`
	EntityName  string         `yaml:"entity_name,omitempty"`
	EntityRef   string         `yaml:"entity_ref,omitempty"`
	Status      PlanStepStatus `yaml:"status"`
	Description string         `yaml:"description,omitempty"`
}

// PlanEntity represents an implementation plan.
type PlanEntity struct {
	BaseEntity    `yaml:",inline"`
	Source        string     `yaml:"source,omitempty"`
	ClaudePlanFile string   `yaml:"claude_plan_file,omitempty"`
	CreatedAt     string     `yaml:"created_at,omitempty"`
	ApprovedAt    string     `yaml:"approved_at,omitempty"`
	CompletedAt   string     `yaml:"completed_at,omitempty"`
	Steps         []PlanStep `yaml:"steps,omitempty"`
}

// Progress returns the completion percentage of the plan.
func (p *PlanEntity) Progress() float64 {
	if len(p.Steps) == 0 {
		return 0
	}
	completed := 0
	for _, s := range p.Steps {
		if s.Status == StepCompleted || s.Status == StepSkipped {
			completed++
		}
	}
	return float64(completed) / float64(len(p.Steps)) * 100
}
