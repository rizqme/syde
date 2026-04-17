---
id: REQ-0350
kind: requirement
name: Each flow step shall reference at most one contract
slug: each-flow-step-shall-reference-at-most-one-contract-5tzk
description: contract field is a single slug, not a list
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The syde entity model shall restrict each flow step to reference at most one contract by slug.
req_type: constraint
priority: must
verification: FlowStep.Contract is string not []string
source: plan
requirement_status: active
rationale: If a behavior touches two contracts, model as two steps
---
