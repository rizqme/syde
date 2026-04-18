---
id: REQ-0214
kind: requirement
name: List Entities Returns Entities Array
slug: list-entities-returns-entities-array-0udf
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:38:05Z"
statement: When a client invokes GET /api/<project>/entities, the syded daemon shall respond with 200 OK and a JSON body containing an entities array of entity summaries.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: The dashboard entity browser lists all entities for a project.
audited_overlaps:
    - slug: http-api-lists-entities-per-project
      distinction: List-Entities-Returns scopes the response payload shape; HTTP-API-Lists scopes the per-project routing dimension — different concerns
    - slug: http-api-lists-entities-per-project-1yjc
      distinction: Specifies 200 OK status and entities array body shape; target specifies per-project scoping of the response.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:38:05Z"
---
