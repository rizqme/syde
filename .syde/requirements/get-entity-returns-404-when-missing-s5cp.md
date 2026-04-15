---
id: REQ-0228
kind: requirement
name: Get Entity Returns 404 When Missing
slug: get-entity-returns-404-when-missing-s5cp
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:41Z"
statement: If the slug supplied to GET /api/<project>/entity/<slug> does not match any entity, then the syded daemon shall respond with HTTP 404.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: Clients must distinguish missing entities from server errors.
---
