---
id: TSK-0048
kind: task
name: Pre-render proposed wireframe HTML for screen contract field_changes
slug: pre-render-proposed-wireframe-html-for-screen-contract-fieldchanges-bts4
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: When an Extended change targets a screen contract and field_changes contains a 'wireframe' key, the API also returns the proposed wireframe rendered to HTML alongside the current one.
details: 'internal/dashboard/api.go handlePlanDetail: when iterating ExtendedChange entries on screen contracts, if e.FieldChanges["wireframe"] is set, parse it through uiml.Parse and call uiml.RenderWireframeHTML to produce a proposed_wireframe_html string. Add it to a new field on the response (e.g. proposed_values_html.wireframe). Pair it with the existing current_values.wireframe_html so the frontend can render both.'
acceptance: curl /api/<proj>/plan/<slug> for a plan that extends a screen contract returns proposed_values_html.wireframe HTML rendered from the new wireframe value.
affected_entities:
    - http-api-afos
affected_files:
    - internal/dashboard/api.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:29:36Z"
completed_at: "2026-04-15T15:07:54Z"
---
