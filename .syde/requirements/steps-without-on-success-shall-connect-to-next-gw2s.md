---
id: REQ-0351
kind: requirement
name: Steps without on-success shall connect to next
slug: steps-without-on-success-shall-connect-to-next-gw2s
description: Empty on_success means implicit next in array order
relationships:
    - target: web-spa
      type: refines
updated_at: "2026-04-18T09:37:44Z"
statement: When a flow step has an empty on_success field, the syde dashboard shall render it as connected to the next step in array order.
req_type: functional
priority: must
verification: Flowchart renders linear steps with implicit edges
source: plan
requirement_status: active
rationale: Convention over configuration for linear flows
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:44Z"
---
