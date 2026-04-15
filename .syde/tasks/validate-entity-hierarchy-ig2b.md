---
acceptance: sync validation reports entities without belongs_to, while allowing only the root system to omit it.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit
    - internal/audit/audit.go
    - internal/audit/graph_rules.go
completed_at: "2026-04-15T06:35:10Z"
created_at: "2026-04-15T06:33:38Z"
details: Add audit graph validation that finds the single root system and reports every other entity missing a belongs_to relationship.
id: TSK-0097
kind: task
name: Validate entity hierarchy
objective: Require every entity except the root system to have a belongs_to parent.
plan_phase: phase_6
plan_ref: add-requirement-entity-56dv
priority: high
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: validate-entity-hierarchy-ig2b
task_status: completed
updated_at: "2026-04-15T06:35:10Z"
---
