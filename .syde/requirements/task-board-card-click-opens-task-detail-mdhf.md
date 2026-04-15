---
id: REQ-0253
kind: requirement
name: Task Board Card Click Opens Task Detail
slug: task-board-card-click-opens-task-detail-mdhf
relationships:
    - target: task-board-screen-tvvy
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:06Z"
statement: When the user clicks a task card on the task board, the dashboard shall open that task's detail view.
req_type: interface
priority: should
verification: manual inspection of /task in the dashboard
source: manual
source_ref: contract:task-board-screen-tvvy
requirement_status: active
rationale: Card click is the expected kanban affordance for inspecting a task.
---
