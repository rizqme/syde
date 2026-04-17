---
id: TSK-0045
kind: task
name: SPA listens for navigate WebSocket messages
slug: spa-listens-for-navigate-websocket-messages-cy9m
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: When the dashboard SPA receives a navigate message over its WebSocket, it programmatically navigates to the supplied path.
details: 'web/src/hooks/useWebSocket.ts (or wherever messages are handled) — add a ''navigate'' case that calls window.history.pushState then dispatches the project''s navigate hook. App.tsx subscribes to navigate events and updates the URL. Idempotent: if the SPA is already on the target path, do nothing.'
acceptance: With the dashboard open in Chrome and a navigate message broadcast for the current project, the visible tab switches to the new URL without reloading.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/hooks/useWebSocket.ts
    - web/src/App.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:26:48Z"
completed_at: "2026-04-15T15:11:22Z"
---
