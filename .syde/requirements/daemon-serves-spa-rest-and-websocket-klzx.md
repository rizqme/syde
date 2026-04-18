---
id: REQ-0026
kind: requirement
name: Daemon Serves SPA REST and WebSocket
slug: daemon-serves-spa-rest-and-websocket-klzx
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
updated_at: "2026-04-18T09:37:26Z"
statement: The syded daemon shall serve the embedded SPA, REST API, and WebSocket endpoints from a single HTTP listener.
req_type: functional
priority: must
verification: integration test hitting /, /api/, and /ws/ on the same port
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: A single binary with one listener simplifies local use and embedding.
verified_against:
    dashboard-daemon-entry-point-qx5c:
        hash: 223060e5cff54830c3871bf36187e234213855226030a492982b9dca51c770cc
        at: "2026-04-18T09:37:26Z"
---
