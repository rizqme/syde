---
id: REQ-0223
kind: requirement
name: Get Entity Returns Full Payload
slug: get-entity-returns-full-payload-fdpe
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:31Z"
statement: When a client invokes GET /api/<project>/entity/<slug>, the syded daemon shall respond with 200 OK and a JSON body containing the full entity object including relationships, files, and body.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: The dashboard detail panel needs the complete entity to render.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:31Z"
---
