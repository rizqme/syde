---
acceptance: go build passes; list --format json returns parseable JSON
completed_at: "2026-04-14T06:19:27Z"
created_at: "2026-04-14T06:04:02Z"
details: 'NEW FILE internal/cli/output.go: Emit(cmd, format, payload). json uses encoding/json; rich calls per-command pretty-printer callback; compact falls through to current path.'
id: TSK-0011
kind: task
name: Create internal/cli/output.go shared Emit formatter
objective: One place owns --format routing (rich/compact/json)
plan_phase: phase_3
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: create-internalclioutputgo-shared-emit-formatter-h05w
task_status: completed
updated_at: "2026-04-14T06:19:27Z"
---
