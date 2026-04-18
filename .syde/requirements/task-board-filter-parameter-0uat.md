---
id: REQ-0251
kind: requirement
name: Task Board Filter Parameter
slug: task-board-filter-parameter-0uat
relationships: []
updated_at: '2026-04-16T01:09:57Z'
statement: Where a filter query parameter is provided on the /task route, the dashboard shall restrict the rendered task columns to tasks matching the status filter.
req_type: interface
priority: should
verification: manual inspection of /task?filter=... in the dashboard
source: manual
source_ref: contract:task-board-screen-tvvy
requirement_status: obsolete
rationale: Status filtering lets users focus the board on a specific slice of work.
obsolete_reason: The standalone task board route was removed; tasks are now reviewed inside plan detail views.
---
