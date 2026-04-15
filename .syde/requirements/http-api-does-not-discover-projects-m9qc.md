---
id: REQ-0049
kind: requirement
name: HTTP API Does Not Discover Projects
slug: http-api-does-not-discover-projects-m9qc
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: The syded HTTP API shall not discover projects on the filesystem by itself.
req_type: constraint
priority: must
verification: inspection confirming project lookup is delegated to Project Registry
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Separating discovery from routing keeps handlers thin and testable.
---
