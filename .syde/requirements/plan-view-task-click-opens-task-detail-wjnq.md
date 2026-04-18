---
id: REQ-0247
kind: requirement
name: Plan View Task Click Opens Task Detail
slug: plan-view-task-click-opens-task-detail-wjnq
relationships:
    - target: plan-view-screen-gb2y
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:38:06Z"
statement: When the user clicks a task row on the plan view screen, the dashboard shall navigate to that task's detail view.
req_type: interface
priority: should
verification: manual inspection of /plan/<slug> in the dashboard
source: manual
source_ref: contract:plan-view-screen-gb2y
requirement_status: active
rationale: Task drill-down is the primary interaction from plan hierarchy to task editing.
audited_overlaps:
    - slug: graph-node-click-navigates-to-detail-8fsr
      distinction: plan-view-task-click triggers within the plan tasks tab; graph-node-click is the Graph canvas detail navigation — different parent screens
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:38:06Z"
---
