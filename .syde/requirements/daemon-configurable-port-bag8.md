---
id: REQ-0025
kind: requirement
name: Daemon Configurable Port
slug: daemon-configurable-port-bag8
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
updated_at: "2026-04-18T09:37:58Z"
statement: The syded daemon shall bind its HTTP listener to a configurable port supplied at launch.
req_type: functional
priority: must
verification: integration test launching syded with a custom --port flag and checking that port
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Operators must be able to avoid port conflicts on developer machines.
verified_against:
    dashboard-daemon-entry-point-qx5c:
        hash: 223060e5cff54830c3871bf36187e234213855226030a492982b9dca51c770cc
        at: "2026-04-18T09:37:58Z"
---
