---
id: TSK-0126
kind: task
name: Wire requirement CLI flags and dashboard fields
slug: wire-requirement-cli-flags-and-dashboard-fields-f9qr
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:38:36Z"
task_status: completed
priority: medium
objective: syde add/update requirement exposes --type --priority --verification --refines flags and the dashboard EntityDetail displays them.
details: 'internal/cli/add.go + update.go: add flag bindings and applyEntityData branches. web/src/components/EntityDetail.tsx: render req_type, priority, verification, refines list for requirement kind entities.'
acceptance: Creating a requirement via CLI stores all four fields; dashboard detail view shows them.
affected_entities:
    - cli-commands-hpjb
    - web-spa-jy9z
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:38:36Z"
---
