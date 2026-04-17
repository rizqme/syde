---
id: REQ-0241
kind: requirement
name: Graph Node Click Navigates To Detail
slug: graph-node-click-navigates-to-detail-8fsr
relationships:
    - target: graph-screen-4044
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-17T11:07:56Z"
statement: When the user clicks a node on the graph canvas, the dashboard shall navigate to that entity's detail view.
req_type: interface
priority: should
verification: manual inspection of /__graph__ in the dashboard
source: manual
source_ref: contract:graph-screen-4044
requirement_status: active
rationale: Node-to-detail navigation bridges graph exploration and entity editing.
audited_overlaps:
    - slug: plan-view-task-click-opens-task-detail-wjnq
      distinction: graph-node-click navigates within the Graph canvas to a detail panel; plan-view-task-click is a distinct surface in the plan tasks tab — different parent screens
---
