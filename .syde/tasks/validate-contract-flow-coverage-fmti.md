---
acceptance: sync validation reports contracts that have no relationship to any flow.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit
    - internal/audit/audit.go
    - internal/audit/graph_rules.go
completed_at: "2026-04-15T06:35:19Z"
created_at: "2026-04-15T06:33:44Z"
details: Add audit graph validation that checks outbound and inbound relationships between contracts and flows.
id: TSK-0098
kind: task
name: Validate contract flow coverage
objective: Require every contract to participate in at least one flow.
plan_phase: phase_7
plan_ref: add-requirement-entity-56dv
priority: high
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: validate-contract-flow-coverage-fmti
task_status: completed
updated_at: "2026-04-15T06:35:19Z"
---
