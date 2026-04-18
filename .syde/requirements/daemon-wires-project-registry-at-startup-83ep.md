---
id: REQ-0028
kind: requirement
name: Daemon Wires Project Registry At Startup
slug: daemon-wires-project-registry-at-startup-83ep
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
updated_at: "2026-04-18T09:36:51Z"
statement: When the syded daemon starts, the daemon shall construct and inject the project registry into HTTP handlers before accepting requests.
req_type: functional
priority: must
verification: integration test asserting /api/projects responds without 500 on first request
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Handlers cannot serve project data without a live registry.
verified_against:
    dashboard-daemon-entry-point-qx5c:
        hash: 223060e5cff54830c3871bf36187e234213855226030a492982b9dca51c770cc
        at: "2026-04-18T09:36:51Z"
---
