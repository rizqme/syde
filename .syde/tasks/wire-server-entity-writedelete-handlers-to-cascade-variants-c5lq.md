---
acceptance: Running syde update component-foo from the CLI bumps both component/foo and system/parent UpdatedAt in one syde task done cycle
affected_entities:
    - http-api
affected_files:
    - internal/dashboard/api_write.go
completed_at: "2026-04-14T07:32:24Z"
created_at: "2026-04-14T07:26:27Z"
details: 'In internal/dashboard/api_write.go: handleEntityWrite switches to store.CreateCascade (POST) and store.UpdateCascade (PUT). handleEntityDelete switches to store.DeleteCascade. No CLI or client changes needed — everything routes through the server and the server now cascades.'
id: TSK-0038
kind: task
name: Wire server entity write/delete handlers to cascade variants
objective: Every HTTP write on /api/<project>/entity goes through the cascade path, so CLI write commands benefit automatically
plan_phase: phase_3
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: wire-server-entity-writedelete-handlers-to-cascade-variants-c5lq
task_status: completed
updated_at: "2026-04-14T07:32:24Z"
---
