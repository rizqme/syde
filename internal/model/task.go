package model

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	TaskPending    TaskStatus = "pending"
	TaskInProgress TaskStatus = "in_progress"
	TaskCompleted  TaskStatus = "completed"
	TaskBlocked    TaskStatus = "blocked"
	TaskCancelled  TaskStatus = "cancelled"
)

// Priority represents task priority.
type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)

// TaskEntity represents a tracked work item.
type TaskEntity struct {
	BaseEntity  `yaml:",inline"`
	TaskStatus  TaskStatus `yaml:"task_status,omitempty"`
	Priority    Priority   `yaml:"priority,omitempty"`
	Objective   string     `yaml:"objective,omitempty"`
	Details     string     `yaml:"details,omitempty"`
	Acceptance  string     `yaml:"acceptance,omitempty"`
	// AffectedEntities are slugs of *existing* entities this task will
	// modify. They are references only — the task does NOT create them.
	// If the task needs to create a brand-new entity, describe it in the
	// task's note field (free text); the agent will run `syde add` for
	// that entity as part of implementation.
	AffectedEntities []string `yaml:"affected_entities,omitempty"`
	// AffectedFiles are concrete source file paths this task will touch.
	// Each entry must exist as a node in the summary tree.
	AffectedFiles []string `yaml:"affected_files,omitempty"`
	PlanRef       string   `yaml:"plan_ref,omitempty"`
	PlanPhase     string   `yaml:"plan_phase,omitempty"`
	// EntityRefs is retained for backwards compatibility with old tasks
	// that used --entity to link tasks to entities. Prefer
	// AffectedEntities on new tasks.
	EntityRefs  []string `yaml:"entity_refs,omitempty"`
	ParentTask  string   `yaml:"parent_task,omitempty"`
	Subtasks    []string `yaml:"subtasks,omitempty"`
	AssignedTo  string   `yaml:"assigned_to,omitempty"`
	BlockReason string   `yaml:"block_reason,omitempty"`
	CreatedAt   string   `yaml:"created_at,omitempty"`
	CompletedAt string   `yaml:"completed_at,omitempty"`
}
