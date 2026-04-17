---
id: TSK-0008
kind: task
name: syde plan remove-change and show-changes
slug: syde-plan-remove-change-and-show-changes-opyq
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: syde plan remove-change <plan> <change-id> deletes the entry; syde plan show-changes <plan> prints a rich summary.
details: remove-change walks all six lanes, finds the entry by id, removes it, rewrites the plan. show-changes renders per-lane headers with Deleted/Extended/New counts and bullets. Support --format json for automated consumption by the dashboard.
acceptance: 'Round-trip: add-change + remove-change leaves the plan byte-identical to pre-add state; show-changes prints every category with non-empty content visible.'
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_2
created_at: "2026-04-15T11:40:57Z"
completed_at: "2026-04-15T11:58:31Z"
---
