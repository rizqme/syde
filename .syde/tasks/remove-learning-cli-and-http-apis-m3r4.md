---
id: TSK-0113
kind: task
name: Remove learning CLI and HTTP APIs
slug: remove-learning-cli-and-http-apis-m3r4
relationships:
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: remove-learning-entity-and-cap-requirement-fanout-wqyz
      type: belongs_to
updated_at: "2026-04-15T09:23:33Z"
task_status: completed
priority: high
objective: Remove syde remember/learn commands and learning endpoints/context payloads.
details: Delete or disconnect internal/cli/remember.go and remove learning handlers and context/constraints learning payloads from dashboard APIs.
acceptance: rg finds no syde remember/learn command registration or /learnings endpoint.
affected_entities:
    - cli-commands-hpjb
    - http-api-afos
affected_files:
    - internal/cli/remember.go
    - internal/dashboard/api.go
    - internal/dashboard/api_readall.go
    - internal/dashboard/html.go
plan_ref: remove-learning-entity-and-cap-requirement-fanout-wqyz
plan_phase: phase_2
created_at: "2026-04-15T08:23:05Z"
completed_at: "2026-04-15T09:04:18Z"
---
