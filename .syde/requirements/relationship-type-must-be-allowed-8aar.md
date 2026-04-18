---
id: REQ-0082
kind: requirement
name: Relationship Type Must Be Allowed
slug: relationship-type-must-be-allowed-8aar
relationships:
    - target: relationship-hjgt
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:08Z"
statement: The syde CLI shall accept a relationship type only if it is one of belongs_to, depends_on, exposes, consumes, uses, involves, references, relates_to, implements, applies_to, modifies, or visualizes.
req_type: constraint
priority: must
verification: integration test adding a relationship with an unknown type
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: A closed type set keeps graph semantics interoperable across tooling.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:08Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:08Z"
---
