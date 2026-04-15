---
acceptance: Creating a screen contract via syde add and then opening the Contracts page shows a Wireframe section in the detail panel with the rendered HTML preview.
affected_entities:
    - query-engine
    - web-spa
affected_files:
    - internal/query/formatter.go
    - web/src/lib/api.ts
    - web/src/components/EntityDetail.tsx
completed_at: "2026-04-14T11:32:17Z"
created_at: "2026-04-14T11:23:14Z"
details: 'internal/query/formatter.go FormatJSON: in the *model.ContractEntity case, if e.ContractKind == "screen" and e.Wireframe != "", parse the UIML and set entityMap[''wireframe'']=raw source, entityMap[''wireframe_html'']=uiml.RenderHTML(nodes), entityMap[''wireframe_ascii'']=uiml.RenderASCII(nodes, 80). web/src/lib/api.ts: contract detail type gains wireframe?: string, wireframe_html?: string, wireframe_ascii?: string. web/src/components/EntityDetail.tsx contract case: add Wireframe section below the existing contract fields, render wireframe_html via dangerouslySetInnerHTML inside a bordered box.'
id: TSK-0077
kind: task
name: Pre-render wireframe HTML/ASCII in FormatJSON + EntityDetail panel
objective: Dashboard shows the UIML wireframe as HTML preview under every screen contract detail
plan_phase: phase_2
plan_ref: screen-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: pre-render-wireframe-htmlascii-in-formatjson-entitydetail-panel-bksw
task_status: completed
updated_at: "2026-04-14T11:32:17Z"
---
