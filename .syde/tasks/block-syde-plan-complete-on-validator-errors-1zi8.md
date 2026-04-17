---
id: TSK-0011
kind: task
name: Block syde plan complete on validator errors
slug: block-syde-plan-complete-on-validator-errors-1zi8
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan complete <slug> refuses to mark the plan completed if planCompletionFindings returns any ERROR for that plan.
details: Add a complete subcommand (or extend the existing execute/done path) that runs the audit focused on the target plan and prints findings. Non-zero exit on ERROR. Use --force to override with a warning.
acceptance: Attempting to complete a plan with a mismatched FieldChanges value fails with the finding printed; --force proceeds with a warning line.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_3
created_at: "2026-04-15T11:41:14Z"
completed_at: "2026-04-15T12:03:26Z"
---
