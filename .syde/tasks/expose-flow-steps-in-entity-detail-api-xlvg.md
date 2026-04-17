---
id: TSK-0061
kind: task
name: Expose flow steps in entity detail API
slug: expose-flow-steps-in-entity-detail-api-xlvg
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: GET /api/<proj>/entity/<slug> returns steps array for flow entities
details: Ensure the entity detail API serializes FlowEntity.Steps as JSON. May already work if the YAML-to-JSON serialization handles it.
acceptance: curl entity/<flow-slug> returns steps array with contract refs
affected_entities:
    - http-api-afos
    - query-engine-9k84
affected_files:
    - internal/dashboard/api.go
    - internal/query/formatter.go
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_3
created_at: "2026-04-16T09:22:48Z"
completed_at: "2026-04-16T10:44:01Z"
---
