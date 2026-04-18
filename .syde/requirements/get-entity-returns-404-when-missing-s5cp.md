---
id: REQ-0228
kind: requirement
name: Get Entity Returns 404 When Missing
slug: get-entity-returns-404-when-missing-s5cp
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:36:49Z"
statement: If the slug supplied to GET /api/<project>/entity/<slug> does not match any entity, then the syded daemon shall respond with HTTP 404.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: Clients must distinguish missing entities from server errors.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:49Z"
---
