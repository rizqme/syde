---
id: TSK-0105
kind: task
name: Wire approval validation before mutation
slug: wire-approval-validation-before-mutation-rita
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:07:40Z"
task_status: completed
priority: high
objective: Call the phase task validation before plan approval mutates status or creates the plan requirement.
details: Wire the helper into planApproveCmd before p.PlanStatus changes, preserving existing requirement capture after validation succeeds.
acceptance: A plan with an empty phase fails before status or requirement changes; a valid plan still creates/reuses the plan requirement.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_2
created_at: "2026-04-15T06:55:24Z"
completed_at: "2026-04-15T07:04:36Z"
---
