---
id: REQ-0272
kind: requirement
name: Get Entity Invocation
slug: get-entity-invocation-il5p
relationships:
    - target: get-entity-0lzq
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:39Z"
statement: When the user runs syde get <slug>, the syde CLI shall render the entity's YAML frontmatter and markdown body on stdout.
req_type: interface
priority: must
verification: integration test invoking syde get on an existing slug
source: manual
source_ref: contract:get-entity-0lzq
requirement_status: active
rationale: syde get is the canonical accessor used by agents and humans alike.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:39Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:39Z"
---
