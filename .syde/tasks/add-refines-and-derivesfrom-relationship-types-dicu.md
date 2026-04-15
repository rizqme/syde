---
id: TSK-0124
kind: task
name: Add refines and derives_from relationship types
slug: add-refines-and-derivesfrom-relationship-types-dicu
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:34:56Z"
task_status: completed
priority: high
objective: Two new relationship type constants exist and are accepted by the relationship parser and audit graph.
details: 'internal/model/relationship.go: add RelRefines and RelDerivesFrom constants. Update any allowed-type whitelists. Document in skill/references/entity-spec.md.'
acceptance: syde update <req> --add-rel other-req:refines succeeds and the target validates.
affected_entities:
    - entity-model-f28o
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:34:56Z"
---
