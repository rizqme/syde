---
boundaries: Does NOT do any business logic itself.
capabilities:
    - Bind to a configurable port
    - Serve the embedded SPA + REST + WebSocket
description: main package for the syded HTTP daemon binary.
files:
    - cmd/syded/main.go
id: COM-0017
kind: component
name: Dashboard Daemon Entry Point
purpose: Boot the syded HTTP daemon from the syded binary
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: depends_on
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Parse flags, start the HTTP server, wire the project registry
slug: dashboard-daemon-entry-point-qx5c
updated_at: "2026-04-14T06:24:45Z"
---
