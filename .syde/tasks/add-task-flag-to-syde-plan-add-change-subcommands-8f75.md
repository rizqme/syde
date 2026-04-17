---
id: TSK-0052
kind: task
name: Add --task flag to syde plan add-change subcommands
slug: add-task-flag-to-syde-plan-add-change-subcommands-8f75
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan add-change delete/extend/new accept a repeatable --task <task-slug> flag that populates the new Tasks field on the appended entry.
details: 'internal/cli/plan.go: add planChangeTasks []string variable and wire it into all three add-change handlers. The flag is repeatable; each occurrence appends one task slug. No validation at this layer — the audit rule (Phase 6) catches empty/unresolvable lists.'
acceptance: syde plan add-change extend <plan> component foo --what x --why y --task task-a --task task-b appends an entry whose tasks list contains [task-a, task-b].
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:38:32Z"
completed_at: "2026-04-15T21:34:15Z"
---
