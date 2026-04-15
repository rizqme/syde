---
id: REQ-0216
kind: requirement
name: List Entities Accepts Kind Filter
slug: list-entities-accepts-kind-filter-rjf6
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:27Z"
statement: When GET /api/<project>/entities is invoked, the syded daemon shall accept an optional kind query parameter as a string and restrict results to that entity kind.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities?kind=component
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: Dashboard pages filter entity lists by kind.
---
