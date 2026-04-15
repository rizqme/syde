---
id: TSK-0106
kind: task
name: Smoke test plan approval task gate
slug: smoke-test-plan-approval-task-gate-73jy
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:07:40Z"
task_status: completed
priority: medium
objective: Verify approval fails for empty phases and succeeds when every phase has tasks.
details: Create or use a draft test plan to exercise the failure path, then verify the updated implementation can approve the real task-covered plan.
acceptance: Terminal output demonstrates the empty-phase failure and successful approval of a task-covered plan.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_2
created_at: "2026-04-15T06:55:31Z"
completed_at: "2026-04-15T07:05:28Z"
---
