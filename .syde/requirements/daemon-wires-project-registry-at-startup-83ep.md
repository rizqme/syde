---
id: REQ-0028
kind: requirement
name: Daemon Wires Project Registry At Startup
slug: daemon-wires-project-registry-at-startup-83ep
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: When the syded daemon starts, the daemon shall construct and inject the project registry into HTTP handlers before accepting requests.
req_type: functional
priority: must
verification: integration test asserting /api/projects responds without 500 on first request
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Handlers cannot serve project data without a live registry.
---
