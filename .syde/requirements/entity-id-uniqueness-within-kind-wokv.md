---
id: REQ-0021
kind: requirement
name: Entity ID Uniqueness Within Kind
slug: entity-id-uniqueness-within-kind-wokv
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:52Z"
statement: The syde CLI shall ensure that every entity ID is unique within its kind and is never reused.
req_type: constraint
priority: must
verification: integration test creating two entities and asserting distinct IDs with no reuse after deletion
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Stable, unique IDs are the backbone of traceability across the syde model.
---
