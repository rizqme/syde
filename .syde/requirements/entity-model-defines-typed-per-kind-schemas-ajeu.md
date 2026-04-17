---
id: REQ-0138
kind: requirement
name: Entity Model Defines Typed Per Kind Schemas
slug: entity-model-defines-typed-per-kind-schemas-ajeu
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:45:32Z"
statement: The entity model shall define typed Go structs for every syde entity kind including system, component, contract, concept, flow, decision, plan, task, design, and requirement.
req_type: functional
priority: must
verification: inspection of internal/model/entity.go per-kind structs
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Typed schemas enforce kind-specific fields at compile time.
audited_overlaps:
    - slug: add-entity-kind-arg-s3jv
      distinction: Defines the internal Go struct typing for every entity kind, not the CLI positional argument validation on syde add.
---
