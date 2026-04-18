---
id: REQ-0140
kind: requirement
name: Entity Model Supports YAML Frontmatter Marshalling
slug: entity-model-supports-yaml-frontmatter-marshalling-6epm
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:15Z"
statement: The entity model shall marshal and unmarshal entities to and from YAML frontmatter using kind-compatible yaml tags.
req_type: functional
priority: must
verification: unit test of marshal/unmarshal round-trip in internal/model
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: YAML frontmatter is the on-disk serialization format for every entity.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:15Z"
---
