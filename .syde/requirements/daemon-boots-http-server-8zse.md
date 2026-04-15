---
id: REQ-0020
kind: requirement
name: Daemon Boots HTTP Server
slug: daemon-boots-http-server-8zse
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:51:48Z"
statement: The syded daemon shall parse flags, start the HTTP server, and wire the project registry on startup.
req_type: functional
priority: must
verification: integration test invoking syded and checking the HTTP port is listening
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Booting the server is the sole reason the entry point exists.
---
