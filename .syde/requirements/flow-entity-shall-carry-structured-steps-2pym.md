---
id: REQ-0346
kind: requirement
name: Flow entity shall carry structured steps
slug: flow-entity-shall-carry-structured-steps-2pym
description: FlowEntity gains a Steps []FlowStep field
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:37:15Z"
statement: The syde entity model shall provide a structured steps field on flow entities.
req_type: functional
priority: must
verification: FlowEntity has Steps []FlowStep with yaml/json tags; round-trip preserves data
source: plan
requirement_status: active
rationale: Structured steps enable programmatic audit and visual rendering
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:15Z"
---
