---
id: TSK-0101
kind: task
name: Enforce tasks during plan approval
slug: enforce-tasks-during-plan-approval-3oqq
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:07:40Z"
task_status: completed
priority: high
objective: Add a focused helper that rejects empty-task phases during plan approval.
details: Implement and unit-smoke the approval-time phase task validation helper in internal/cli/plan.go before any approval state mutation.
acceptance: The helper returns a clear error listing every phase with zero direct tasks.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_2
created_at: "2026-04-15T06:51:44Z"
completed_at: "2026-04-15T07:04:15Z"
---
