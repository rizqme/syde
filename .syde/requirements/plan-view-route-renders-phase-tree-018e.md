---
id: REQ-0242
kind: requirement
name: Plan View Route Renders Phase Tree
slug: plan-view-route-renders-phase-tree-018e
relationships:
    - target: plan-view-screen-gb2y
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:36:50Z"
statement: When the user navigates to the /plan/<slug> route, the dashboard shall render the plan's phases and their tasks with per-task status indicators.
req_type: interface
priority: must
verification: manual inspection of /plan/<slug> in the dashboard
source: manual
source_ref: contract:plan-view-screen-gb2y
requirement_status: active
rationale: Plan view is the hierarchical progress surface for planned work.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:50Z"
---
