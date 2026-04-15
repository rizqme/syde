---
id: REQ-0139
kind: requirement
name: Entity Model Dispatches By Kind
slug: entity-model-dispatches-by-kind-w07l
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:22Z"
statement: When NewEntityForKind is called with a valid EntityKind, the entity model shall return a typed zero-valued entity of that kind.
req_type: functional
priority: must
verification: unit test of NewEntityForKind in internal/model/entity.go
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Kind dispatch is the single entry point for creating new entities.
---
