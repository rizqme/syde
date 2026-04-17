---
id: REQ-0043
kind: requirement
name: HTTP API Lists Entities Per Project
slug: http-api-lists-entities-per-project-1yjc
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-17T10:50:15Z"
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
---
