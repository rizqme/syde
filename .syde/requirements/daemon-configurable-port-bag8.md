---
id: REQ-0025
kind: requirement
name: Daemon Configurable Port
slug: daemon-configurable-port-bag8
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: The syded daemon shall bind its HTTP listener to a configurable port supplied at launch.
req_type: functional
priority: must
verification: integration test launching syded with a custom --port flag and checking that port
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Operators must be able to avoid port conflicts on developer machines.
---
