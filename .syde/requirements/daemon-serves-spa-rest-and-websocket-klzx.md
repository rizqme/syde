---
id: REQ-0026
kind: requirement
name: Daemon Serves SPA REST and WebSocket
slug: daemon-serves-spa-rest-and-websocket-klzx
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: The syded daemon shall serve the embedded SPA, REST API, and WebSocket endpoints from a single HTTP listener.
req_type: functional
priority: must
verification: integration test hitting /, /api/, and /ws/ on the same port
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: A single binary with one listener simplifies local use and embedding.
---
