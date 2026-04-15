---
acceptance: syde sync check --strict still returns exit 1/2/0 as before, now via HTTP
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/validate.go
    - internal/cli/sync_check.go
    - internal/cli/constraints.go
    - internal/cli/files.go
completed_at: "2026-04-14T06:53:18Z"
created_at: "2026-04-14T06:38:16Z"
details: Rewire validate.go and sync_check.go to call client.SyncCheck and render the same grouped output from reportPayload. constraints check + files orphans/coverage use client endpoints.
id: TSK-0023
kind: task
name: Rewire health CLI commands (validate, sync check, constraints, files)
objective: Health commands call /api/<project>/validate or /sync-check and print the JSON findings
plan_phase: phase_4
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-health-cli-commands-validate-sync-check-constraints-files-0x4t
task_status: completed
updated_at: "2026-04-14T06:53:18Z"
---
