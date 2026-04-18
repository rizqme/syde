---
id: REQ-0335
kind: requirement
name: Plan changes view shall render screen wireframes inline
slug: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
description: Plan changes wireframe rendering requirement for screen contracts.
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: uiml-parser-sjdk
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:06Z"
statement: When the plan changes view renders an Extended or New change targeting a screen-kind contract, the syded dashboard shall render the contract's wireframe as HTML using the canonical UIML renderer.
req_type: usability
priority: must
verification: 'Manual inspection: open a plan whose changes block extends a screen contract, confirm the wireframe renders visually inside the Extended card.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: UIML source is unreadable to humans during review; rendered wireframes are the whole point of having a wireframe field.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:06Z"
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:37:06Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:06Z"
---
