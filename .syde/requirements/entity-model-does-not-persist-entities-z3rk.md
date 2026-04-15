---
id: REQ-0144
kind: requirement
name: Entity Model Does Not Persist Entities
slug: entity-model-does-not-persist-entities-z3rk
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:33Z"
statement: The entity model shall not persist entities to disk and shall delegate storage to the storage engine.
req_type: constraint
priority: must
verification: code review of internal/model for filesystem imports
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Separating schema from storage keeps validation pure and fast.
---
