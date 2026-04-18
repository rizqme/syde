---
id: REQ-0204
kind: requirement
name: Overview Route Renders Entity Counts
slug: overview-route-renders-entity-counts-ipc8
relationships:
    - target: overview-screen-2011
      type: refines
    - target: web-spa-jy9z
      type: refines
    - target: syded-dashboard-e82c
      type: relates_to
updated_at: "2026-04-18T09:36:51Z"
statement: When the user navigates to the / route, the dashboard shall render an entity counts grid and a recent activity feed.
req_type: interface
priority: must
verification: manual inspection of / in the dashboard
source: manual
source_ref: contract:overview-screen-2011
requirement_status: active
rationale: Overview is the project home and must summarize model size at a glance.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:51Z"
---
