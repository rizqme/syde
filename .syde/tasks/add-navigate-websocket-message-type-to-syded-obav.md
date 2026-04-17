---
id: TSK-0044
kind: task
name: Add navigate WebSocket message type to syded
slug: add-navigate-websocket-message-type-to-syded-obav
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syded broadcasts a navigate event to all connected dashboard clients of a project; clients receive {type:'navigate', path:'/<proj>/plan/<slug>'} and switch their URL to it.
details: 'internal/dashboard/websocket.go: extend the existing per-project WebSocket broadcaster with a NavigateAll(projectSlug, path) method. The message JSON shape is {"type":"navigate","path":"..."}. internal/dashboard/api.go: add POST /api/<proj>/navigate handler that takes {"path":"..."} body and calls NavigateAll(projectSlug, path). Returns 200 with {"clients": <count>} so the caller knows whether anyone was listening.'
acceptance: curl -X POST /api/syde-3646/navigate -d '{"path":"/syde-3646/plan/foo"}' broadcasts to connected clients and returns the listener count.
affected_entities:
    - websocket-server-hdup
    - http-api-afos
affected_files:
    - internal/dashboard/websocket.go
    - internal/dashboard/api.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:26:48Z"
completed_at: "2026-04-15T15:10:50Z"
---
