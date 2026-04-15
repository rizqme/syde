---
id: REQ-0143
kind: requirement
name: Entity Model Requires Description On Every Entity
slug: entity-model-requires-description-on-every-entity-6kx9
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:31Z"
statement: When ValidateEntity is invoked on any entity kind, the entity model shall require a non-empty description field.
req_type: constraint
priority: must
verification: unit test asserting empty description fails ValidateEntity
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Description is the minimum human-readable context needed in every design node.
---
