---
id: TSK-0005
kind: task
name: syde plan add-change delete subcommand
slug: syde-plan-add-change-delete-subcommand-619a
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan add-change <plan> <kind> delete <slug> --why appends a DeletedChange to the correct lane.
details: Add a cobra sub-tree 'add-change' under plan with per-kind and per-op cobra commands. Validate the kind against ValidEntityKinds. Allocate a new change id. Call store.Update on the plan with the appended entry.
acceptance: After running the command, syde query <plan-slug> --full shows the new DeletedChange entry.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_2
created_at: "2026-04-15T11:40:57Z"
completed_at: "2026-04-15T11:52:53Z"
---
