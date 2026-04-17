---
id: REQ-0347
kind: requirement
name: Flow step shall have six fields
slug: flow-step-shall-have-six-fields-znh5
description: FlowStep struct has id, action, contract, description, on_success, on_failure
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The syde entity model shall define each flow step with an id, action, contract, description, on_success, and on_failure field.
req_type: functional
priority: must
verification: FlowStep struct has all six fields with correct yaml/json tags
source: plan
requirement_status: active
rationale: 'Each field serves a distinct purpose: id for linking, action for display, contract for traceability, on_success/on_failure for graph edges'
---
