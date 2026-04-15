---
id: REQ-0029
kind: requirement
name: Daemon Fails Fast On Port Conflict
slug: daemon-fails-fast-on-port-conflict-nvf3
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: If the configured port is already in use, then the syded daemon shall exit with a non-zero status and a clear error message.
req_type: functional
priority: must
verification: integration test launching two syded instances on the same port
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Silent port fallback would confuse operators about which instance is live.
---
