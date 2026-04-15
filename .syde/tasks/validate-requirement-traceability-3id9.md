---
acceptance: sync validation reports non-requirement entities with no relationship to a requirement.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit
    - internal/audit/audit.go
    - internal/audit/graph_rules.go
completed_at: "2026-04-15T06:35:01Z"
created_at: "2026-04-15T06:33:31Z"
details: Add audit graph validation that resolves relationship targets both outbound and inbound, then reports entities without a requirement edge.
id: TSK-0096
kind: task
name: Validate requirement traceability
objective: Require every non-requirement design entity to link to at least one requirement.
plan_phase: phase_5
plan_ref: add-requirement-entity-56dv
priority: high
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: validate-requirement-traceability-3id9
task_status: completed
updated_at: "2026-04-15T06:35:01Z"
---
