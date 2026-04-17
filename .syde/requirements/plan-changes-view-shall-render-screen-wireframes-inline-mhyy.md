---
id: REQ-0335
kind: requirement
name: Plan changes view shall render screen wireframes inline
slug: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
description: Plan changes wireframe rendering requirement for screen contracts.
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-16T01:09:24Z"
statement: When the plan changes view renders an Extended or New change targeting a screen-kind contract, the syded dashboard shall render the contract's wireframe as HTML using the canonical UIML renderer.
req_type: usability
priority: must
verification: 'Manual inspection: open a plan whose changes block extends a screen contract, confirm the wireframe renders visually inside the Extended card.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: UIML source is unreadable to humans during review; rendered wireframes are the whole point of having a wireframe field.
---
