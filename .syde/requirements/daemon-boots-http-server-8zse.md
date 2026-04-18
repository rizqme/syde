---
id: REQ-0020
kind: requirement
name: Daemon Boots HTTP Server
slug: daemon-boots-http-server-8zse
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
updated_at: "2026-04-18T09:37:24Z"
statement: The syded daemon shall parse flags, start the HTTP server, and wire the project registry on startup.
req_type: functional
priority: must
verification: integration test invoking syded and checking the HTTP port is listening
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Booting the server is the sole reason the entry point exists.
verified_against:
    dashboard-daemon-entry-point-qx5c:
        hash: 223060e5cff54830c3871bf36187e234213855226030a492982b9dca51c770cc
        at: "2026-04-18T09:37:24Z"
---
