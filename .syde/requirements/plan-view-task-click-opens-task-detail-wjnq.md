---
id: REQ-0247
kind: requirement
name: Plan View Task Click Opens Task Detail
slug: plan-view-task-click-opens-task-detail-wjnq
relationships:
    - target: plan-view-screen-gb2y
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:01Z"
statement: When the user clicks a task row on the plan view screen, the dashboard shall navigate to that task's detail view.
req_type: interface
priority: should
verification: manual inspection of /plan/<slug> in the dashboard
source: manual
source_ref: contract:plan-view-screen-gb2y
requirement_status: active
rationale: Task drill-down is the primary interaction from plan hierarchy to task editing.
---
