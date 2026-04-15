---
id: REQ-0068
kind: requirement
name: Registry Holds Handles For Daemon Lifetime
slug: registry-holds-handles-for-daemon-lifetime-9zx7
relationships:
    - target: project-registry-q1zx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: While the syded daemon is running, the project registry shall hold each cached Store handle open and shall not close it on caller request.
req_type: constraint
priority: must
verification: inspection confirming GetStore returns a handle without a Close method exposed to callers
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Closing a shared handle would break concurrent handlers mid-request.
---
