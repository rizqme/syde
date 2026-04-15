---
id: REQ-0214
kind: requirement
name: List Entities Returns Entities Array
slug: list-entities-returns-entities-array-0udf
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:24Z"
statement: When a client invokes GET /api/<project>/entities, the syded daemon shall respond with 200 OK and a JSON body containing an entities array of entity summaries.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: The dashboard entity browser lists all entities for a project.
---
