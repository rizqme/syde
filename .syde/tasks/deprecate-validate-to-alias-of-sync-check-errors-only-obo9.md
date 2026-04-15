---
acceptance: syde validate still works but shows deprecation hint
affected_files:
    - internal/cli/validate.go
completed_at: "2026-04-14T06:17:13Z"
created_at: "2026-04-14T06:03:22Z"
details: validate prints deprecation notice on stderr; forwards to sync check with --errors-only filter. Update skill/references/commands.md.
id: TSK-0002
kind: task
name: Deprecate validate to alias of sync check --errors-only
objective: validate keeps working but steers users to the canonical command
plan_phase: phase_2
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: deprecate-validate-to-alias-of-sync-check-errors-only-obo9
task_status: completed
updated_at: "2026-04-14T06:17:13Z"
---
