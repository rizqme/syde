---
id: REQ-0021
kind: requirement
name: Entity ID Uniqueness Within Kind
slug: entity-id-uniqueness-within-kind-wokv
relationships:
    - target: entity-8x6p
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:00Z"
statement: The syde CLI shall ensure that every entity ID is unique within its kind and is never reused.
req_type: constraint
priority: must
verification: integration test creating two entities and asserting distinct IDs with no reuse after deletion
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Stable, unique IDs are the backbone of traceability across the syde model.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:00Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:00Z"
---
