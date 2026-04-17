---
id: REQ-0351
kind: requirement
name: Steps without on-success shall connect to next
slug: steps-without-on-success-shall-connect-to-next-gw2s
description: Empty on_success means implicit next in array order
relationships:
    - target: syde
      type: belongs_to
    - target: web-spa
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: When a flow step has an empty on_success field, the syde dashboard shall render it as connected to the next step in array order.
req_type: functional
priority: must
verification: Flowchart renders linear steps with implicit edges
source: plan
requirement_status: active
rationale: Convention over configuration for linear flows
---
