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
	Priority    Priority `yaml:"priority,omitempty"`
	PlanRef     string   `yaml:"plan_ref,omitempty"`
	PlanStep    string   `yaml:"plan_step,omitempty"`
	EntityRefs  []string `yaml:"entity_refs,omitempty"`
	ParentTask  string   `yaml:"parent_task,omitempty"`
	Subtasks    []string `yaml:"subtasks,omitempty"`
	AssignedTo  string   `yaml:"assigned_to,omitempty"`
	BlockReason string   `yaml:"block_reason,omitempty"`
	CreatedAt   string   `yaml:"created_at,omitempty"`
	CompletedAt string   `yaml:"completed_at,omitempty"`
}
