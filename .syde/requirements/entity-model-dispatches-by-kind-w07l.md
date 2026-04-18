---
id: REQ-0139
kind: requirement
name: Entity Model Dispatches By Kind
slug: entity-model-dispatches-by-kind-w07l
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:04Z"
statement: When NewEntityForKind is called with a valid EntityKind, the entity model shall return a typed zero-valued entity of that kind.
req_type: functional
priority: must
verification: unit test of NewEntityForKind in internal/model/entity.go
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Kind dispatch is the single entry point for creating new entities.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:04Z"
---
