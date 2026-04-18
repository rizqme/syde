---
id: REQ-0238
kind: requirement
name: Graph Route Renders Relationship Canvas
slug: graph-route-renders-relationship-canvas-io70
relationships:
    - target: graph-screen-4044
      type: refines
    - target: web-spa-jy9z
      type: refines
    - target: graph-engine-xgjy
      type: refines
updated_at: "2026-04-18T09:37:06Z"
statement: When the user navigates to the /__graph__ route, the dashboard shall render a relationship graph canvas with reset and fit controls.
req_type: interface
priority: must
verification: manual inspection of /__graph__ in the dashboard
source: manual
source_ref: contract:graph-screen-4044
requirement_status: active
rationale: Graph screen provides the cross-kind relationship overview.
verified_against:
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:37:06Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:06Z"
---
