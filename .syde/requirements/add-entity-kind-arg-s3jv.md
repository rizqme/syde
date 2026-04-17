---
id: REQ-0243
kind: requirement
name: Add Entity Kind Arg
slug: add-entity-kind-arg-s3jv
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-17T10:45:19Z"
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
---
