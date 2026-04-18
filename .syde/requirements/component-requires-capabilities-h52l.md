---
id: REQ-0032
kind: requirement
name: Component Requires Capabilities
slug: component-requires-capabilities-h52l
relationships:
    - target: entity-8x6p
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:43Z"
statement: The syde CLI shall require a purpose, a responsibility, and at least one capability on every component entity instance.
req_type: constraint
priority: must
verification: integration test running syde add component without capabilities
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Components without stated capabilities cannot drive requirement derivation.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:43Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:43Z"
---
