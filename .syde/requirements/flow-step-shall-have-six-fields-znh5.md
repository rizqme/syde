---
id: REQ-0347
kind: requirement
name: Flow step shall have six fields
slug: flow-step-shall-have-six-fields-znh5
description: FlowStep struct has id, action, contract, description, on_success, on_failure
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:36:41Z"
statement: The syde entity model shall define each flow step with an id, action, contract, description, on_success, and on_failure field.
req_type: functional
priority: must
verification: FlowStep struct has all six fields with correct yaml/json tags
source: plan
requirement_status: active
rationale: 'Each field serves a distinct purpose: id for linking, action for display, contract for traceability, on_success/on_failure for graph edges'
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:41Z"
---
