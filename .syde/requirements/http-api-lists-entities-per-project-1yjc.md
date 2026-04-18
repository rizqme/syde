---
id: REQ-0043
kind: requirement
name: HTTP API Lists Entities Per Project
slug: http-api-lists-entities-per-project-1yjc
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:44Z"
statement: When a client invokes GET /api/<project>/entities, the syded daemon shall respond with a JSON list of entities scoped to that project.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The SPA depends on per-project entity listings for every inbox view.
audited_overlaps:
    - slug: list-entities-returns-entities-array
      distinction: HTTP-API-Lists scopes the per-project routing dimension; List-Entities-Returns scopes the response payload shape — different concerns
    - slug: list-entities-returns-entities-array-0udf
      distinction: Specifies per-project scoping of the response; target specifies the 200 OK status and entities array shape.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:44Z"
---
