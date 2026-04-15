---
id: REQ-0238
kind: requirement
name: Graph Route Renders Relationship Canvas
slug: graph-route-renders-relationship-canvas-io70
relationships:
    - target: graph-screen-4044
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:52Z"
statement: When the user navigates to the /__graph__ route, the dashboard shall render a relationship graph canvas with reset and fit controls.
req_type: interface
priority: must
verification: manual inspection of /__graph__ in the dashboard
source: manual
source_ref: contract:graph-screen-4044
requirement_status: active
rationale: Graph screen provides the cross-kind relationship overview.
---
