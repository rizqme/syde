---
id: REQ-0225
kind: requirement
name: Get Entity Accepts Full Or Bare Slug
slug: get-entity-accepts-full-or-bare-slug-qh8d
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:36:52Z"
statement: When GET /api/<project>/entity/<slug> is invoked, the syded daemon shall accept the slug path parameter as a string in either full or bare form and resolve it to a single entity.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: Callers may pass either the full slug or the bare name portion.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:52Z"
---
