---
id: REQ-0346
kind: requirement
name: Flow entity shall carry structured steps
slug: flow-entity-shall-carry-structured-steps-2pym
description: FlowEntity gains a Steps []FlowStep field
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The syde entity model shall provide a structured steps field on flow entities.
req_type: functional
priority: must
verification: FlowEntity has Steps []FlowStep with yaml/json tags; round-trip preserves data
source: plan
requirement_status: active
rationale: Structured steps enable programmatic audit and visual rendering
---
