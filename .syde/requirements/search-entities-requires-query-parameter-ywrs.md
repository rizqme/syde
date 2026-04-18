---
id: REQ-0233
kind: requirement
name: Search Entities Requires Query Parameter
slug: search-entities-requires-query-parameter-ywrs
relationships:
    - target: search-entities-http-1th3
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:20Z"
statement: When GET /api/<project>/search is invoked, the syded daemon shall require a non-empty q query parameter as a string.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: contract:search-entities-http-1th3
requirement_status: active
rationale: Search without a query has no defined meaning.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:20Z"
---
