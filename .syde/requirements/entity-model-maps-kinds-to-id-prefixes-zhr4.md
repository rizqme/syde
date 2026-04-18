---
id: REQ-0142
kind: requirement
name: Entity Model Maps Kinds To ID Prefixes
slug: entity-model-maps-kinds-to-id-prefixes-zhr4
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:29Z"
statement: The entity model shall map every EntityKind to a fixed three-character ID prefix used in counter-based IDs.
req_type: functional
priority: must
verification: inspection of ID prefix mapping in entity.go
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Stable ID prefixes enable ID parsing and cross-references.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:29Z"
---
