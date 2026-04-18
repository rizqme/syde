---
id: REQ-0138
kind: requirement
name: Entity Model Defines Typed Per Kind Schemas
slug: entity-model-defines-typed-per-kind-schemas-ajeu
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:27Z"
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
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:27Z"
---
