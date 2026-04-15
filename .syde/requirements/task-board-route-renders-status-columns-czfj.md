---
id: REQ-0249
kind: requirement
name: Task Board Route Renders Status Columns
slug: task-board-route-renders-status-columns-czfj
relationships:
    - target: task-board-screen-tvvy
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:03Z"
statement: When the user navigates to the /task route, the dashboard shall render a kanban board grouping tasks into Pending, In Progress, and Done columns.
req_type: interface
priority: must
verification: manual inspection of /task in the dashboard
source: manual
source_ref: contract:task-board-screen-tvvy
requirement_status: active
rationale: Task board is the primary kanban surface for tracking work in progress.
---
