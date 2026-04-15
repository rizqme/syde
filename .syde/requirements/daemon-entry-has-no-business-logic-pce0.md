---
id: REQ-0027
kind: requirement
name: Daemon Entry Has No Business Logic
slug: daemon-entry-has-no-business-logic-pce0
relationships:
    - target: dashboard-daemon-entry-point-qx5c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: The syded daemon entry point shall not contain business logic beyond flag parsing and server wiring.
req_type: constraint
priority: must
verification: inspection of cmd/syded/main.go for delegation-only structure
source: manual
source_ref: component:dashboard-daemon-entry-point-qx5c
requirement_status: active
rationale: Keeping main thin makes the daemon easy to test and refactor.
---
