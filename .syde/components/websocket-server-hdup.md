---
boundaries: Does NOT accept writes from clients. Does NOT replace the REST API for bulk fetches.
capabilities:
    - Per-project WebSocket rooms
    - File watcher integration for live refresh
description: Per-project WebSocket broadcaster pushing live updates to the SPA.
files:
    - internal/dashboard/websocket.go
id: COM-0016
kind: component
name: WebSocket Server
purpose: Push live updates to the dashboard SPA when files change on disk
relationships:
    - target: syded-dashboard
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Broadcast entity/tree change events over WebSocket
slug: websocket-server-hdup
updated_at: "2026-04-14T03:35:55Z"
---
