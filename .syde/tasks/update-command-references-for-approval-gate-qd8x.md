---
id: TSK-0103
kind: task
name: Update command references for approval gate
slug: update-command-references-for-approval-gate-qd8x
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:07:40Z"
task_status: completed
priority: medium
objective: Document that plan approval requires task coverage in every phase.
details: Update skill/references/commands.md with the plan approve task coverage requirement and examples that create phase tasks before approval.
acceptance: Command reference says syde plan approve rejects phases without tasks.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/references/commands.md
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_1
created_at: "2026-04-15T06:55:07Z"
completed_at: "2026-04-15T07:01:58Z"
---
