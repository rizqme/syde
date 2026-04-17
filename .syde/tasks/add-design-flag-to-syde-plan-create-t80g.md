---
id: TSK-0004
kind: task
name: Add --design flag to syde plan create
slug: add-design-flag-to-syde-plan-create-t80g
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan create accepts --design and stores it on the new plan.
details: Extend internal/cli/plan.go plan create command with StringVar(&planDesign, 'design', '', 'detailed implementation design prose'). Wire into the PlanEntity constructed by the handler. Update CLI help text.
acceptance: 'syde plan create ''X'' --design ''Y'' creates a plan file whose yaml frontmatter has design: Y.'
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_2
created_at: "2026-04-15T11:40:57Z"
completed_at: "2026-04-15T11:50:40Z"
---
