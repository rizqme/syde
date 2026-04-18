---
id: REQ-0350
kind: requirement
name: Each flow step shall reference at most one contract
slug: each-flow-step-shall-reference-at-most-one-contract-5tzk
description: contract field is a single slug, not a list
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:36:41Z"
statement: The syde entity model shall restrict each flow step to reference at most one contract by slug.
req_type: constraint
priority: must
verification: FlowStep.Contract is string not []string
source: plan
requirement_status: active
rationale: If a behavior touches two contracts, model as two steps
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:41Z"
---
