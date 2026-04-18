---
id: REQ-0243
kind: requirement
name: Add Entity Kind Arg
slug: add-entity-kind-arg-s3jv
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:07Z"
statement: When syde add is invoked, the syde CLI shall require a positional kind argument that is one of system, component, contract, concept, flow, decision, plan, task, design, or requirement.
req_type: interface
priority: must
verification: integration test invoking syde add with each kind
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: Kind dispatch determines which entity schema applies and which fields are mandatory.
audited_overlaps:
    - slug: entity-model-defines-typed-per-kind-schemas-ajeu
      distinction: CLI positional arg validation on syde add at runtime vs. Go type definitions for entity structs at build time — different layer (CLI parsing vs model package).
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:07Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:38:07Z"
---
