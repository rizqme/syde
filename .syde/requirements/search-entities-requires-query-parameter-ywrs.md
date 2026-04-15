---
id: REQ-0233
kind: requirement
name: Search Entities Requires Query Parameter
slug: search-entities-requires-query-parameter-ywrs
relationships:
    - target: search-entities-http-1th3
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:47Z"
statement: When GET /api/<project>/search is invoked, the syded daemon shall require a non-empty q query parameter as a string.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: contract:search-entities-http-1th3
requirement_status: active
rationale: Search without a query has no defined meaning.
---
