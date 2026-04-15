---
acceptance: syde list component --format json returns array; syde query cli-commands --format json includes every field rich shows
affected_files:
    - internal/cli/list.go
    - internal/cli/query.go
    - internal/cli/get.go
    - internal/cli/status.go
completed_at: "2026-04-14T06:20:36Z"
created_at: "2026-04-14T06:03:22Z"
details: Wire --format flag on list (currently compact-only). query/get already have rich/compact — add json. status currently text-only — add json with counts map.
id: TSK-0004
kind: task
name: Add --format json to list, query, get, status
objective: All read commands expose lossless JSON
plan_phase: phase_3
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-format-json-to-list-query-get-status-v8vi
task_status: completed
updated_at: "2026-04-14T06:20:36Z"
---
