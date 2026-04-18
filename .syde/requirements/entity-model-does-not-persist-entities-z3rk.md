---
id: REQ-0144
kind: requirement
name: Entity Model Does Not Persist Entities
slug: entity-model-does-not-persist-entities-z3rk
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:33Z"
statement: The entity model shall not persist entities to disk and shall delegate storage to the storage engine.
req_type: constraint
priority: must
verification: code review of internal/model for filesystem imports
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Separating schema from storage keeps validation pure and fast.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:33Z"
---
