---
id: REQ-0223
kind: requirement
name: Get Entity Returns Full Payload
slug: get-entity-returns-full-payload-fdpe
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:36Z"
statement: When a client invokes GET /api/<project>/entity/<slug>, the syded daemon shall respond with 200 OK and a JSON body containing the full entity object including relationships, files, and body.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: The dashboard detail panel needs the complete entity to render.
---
