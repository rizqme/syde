---
id: REQ-0241
kind: requirement
name: Graph Node Click Navigates To Detail
slug: graph-node-click-navigates-to-detail-8fsr
relationships:
    - target: graph-screen-4044
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:36:53Z"
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
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:53Z"
---
