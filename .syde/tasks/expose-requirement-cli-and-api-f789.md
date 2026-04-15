---
acceptance: syde add requirement can create a requirement; syde update can set requirement fields/statuses; syde list/get/query/status include requirements; syde remove requirement refuses deletion.
affected_entities:
    - cli-commands-hpjb
    - http-api-afos
    - query-engine-9k84
affected_files:
    - internal/cli/add.go
    - internal/cli/update.go
    - internal/cli/remove.go
    - internal/query/formatter.go
    - internal/dashboard/api.go
    - internal/dashboard/api_readall.go
    - internal/query/engine.go
    - internal/client/client.go
    - internal/dashboard/api_write.go
completed_at: "2026-04-15T06:26:47Z"
created_at: "2026-04-15T06:21:27Z"
details: Add requirement-specific add/update flags and formatting, block requirement deletion, and ensure requirements appear in query/status/context/dashboard API payloads.
id: TSK-0093
kind: task
name: Expose requirement CLI and API
objective: Expose requirement entities through CLI read/write paths and daemon API output.
plan_phase: phase_2
plan_ref: add-requirement-entity-56dv
priority: high
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: expose-requirement-cli-and-api-f789
task_status: completed
updated_at: "2026-04-15T06:26:47Z"
---
