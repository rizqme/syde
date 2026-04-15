---
id: COM-0016
kind: component
name: WebSocket Server
slug: websocket-server-hdup
description: Per-project WebSocket broadcaster pushing live updates to the SPA.
purpose: Push live updates to the dashboard SPA when files change on disk
files:
    - internal/dashboard/websocket.go
relationships:
    - target: syded-dashboard
      type: belongs_to
updated_at: "2026-04-15T09:27:08Z"
responsibility: Broadcast entity/tree change events over WebSocket
capabilities:
    - Per-project WebSocket rooms
    - File watcher integration for live refresh
boundaries: Does NOT accept writes from clients. Does NOT replace the REST API for bulk fetches.
---
