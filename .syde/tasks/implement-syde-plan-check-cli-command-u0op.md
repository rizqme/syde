---
id: TSK-0038
kind: task
name: Implement syde plan check CLI command
slug: implement-syde-plan-check-cli-command-u0op
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan check <plan-slug> filters audit findings to plan_authoring + plan_completion scoped to the plan and prints them grouped by severity, exiting non-zero on ERROR.
details: 'internal/cli/plan.go: new planCheckCmd similar to planCompleteCmd but printing both rule sets and not mutating the plan. Exits 0 if zero ERROR findings; non-zero on any ERROR. WARN findings print but don''t fail the gate. Wire into the planCmd subtree.'
acceptance: syde plan check plans-inbox-2-column-layout exits 0 after the plan's gaps are addressed.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_6
created_at: "2026-04-15T13:15:56Z"
completed_at: "2026-04-15T21:39:15Z"
---
