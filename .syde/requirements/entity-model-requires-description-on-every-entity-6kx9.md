---
id: REQ-0143
kind: requirement
name: Entity Model Requires Description On Every Entity
slug: entity-model-requires-description-on-every-entity-6kx9
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:36Z"
statement: When ValidateEntity is invoked on any entity kind, the entity model shall require a non-empty description field.
req_type: constraint
priority: must
verification: unit test asserting empty description fails ValidateEntity
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Description is the minimum human-readable context needed in every design node.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:36Z"
---
