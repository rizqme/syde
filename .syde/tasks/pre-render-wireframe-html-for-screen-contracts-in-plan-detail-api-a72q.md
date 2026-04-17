---
id: TSK-0034
kind: task
name: Pre-render wireframe HTML for screen contracts in plan detail API
slug: pre-render-wireframe-html-for-screen-contracts-in-plan-detail-api-a72q
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: GET /api/<proj>/plan/<slug> returns wireframe_html in current_values for any Extended change targeting a screen-kind contract, and in draft for any New change with contract_kind=screen.
details: 'internal/dashboard/api.go handlePlanDetail: when iterating Extended changes, if the target contract is contract_kind=screen and has Wireframe text, parse via uiml.Parse and call uiml.RenderWireframeHTML, then add the rendered HTML under current_values[''wireframe_html'']. When iterating New changes, if draft.contract_kind === ''screen'' and draft.wireframe is a non-empty string, do the same and add draft[''wireframe_html'']. The frontend can then render via dangerouslySetInnerHTML the same way EntityDetail already does for screen contracts.'
acceptance: curl /api/<proj>/plan/plans-inbox-2-column-layout returns components/contracts.extended[0].current_values.wireframe_html populated for the plan-view-screen Extended entry.
affected_entities:
    - http-api-afos
affected_files:
    - internal/dashboard/api.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:07:14Z"
completed_at: "2026-04-15T15:07:48Z"
---
