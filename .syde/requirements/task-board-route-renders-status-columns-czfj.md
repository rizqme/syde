---
id: REQ-0249
kind: requirement
name: Task Board Route Renders Status Columns
slug: task-board-route-renders-status-columns-czfj
relationships:
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-16T01:09:57Z"
statement: When the user navigates to the /task route, the dashboard shall render a kanban board grouping tasks into Pending, In Progress, and Done columns.
req_type: interface
priority: must
verification: manual inspection of /task in the dashboard
source: manual
source_ref: contract:task-board-screen-tvvy
requirement_status: obsolete
rationale: Task board is the primary kanban surface for tracking work in progress.
obsolete_reason: The standalone task board route was removed; tasks are now reviewed inside plan detail views.
---
