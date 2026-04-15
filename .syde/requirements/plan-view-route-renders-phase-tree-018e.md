---
id: REQ-0242
kind: requirement
name: Plan View Route Renders Phase Tree
slug: plan-view-route-renders-phase-tree-018e
relationships:
    - target: plan-view-screen-gb2y
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:59Z"
statement: When the user navigates to the /plan/<slug> route, the dashboard shall render the plan's phases and their tasks with per-task status indicators.
req_type: interface
priority: must
verification: manual inspection of /plan/<slug> in the dashboard
source: manual
source_ref: contract:plan-view-screen-gb2y
requirement_status: active
rationale: Plan view is the hierarchical progress surface for planned work.
---
