---
id: REQ-0230
kind: requirement
name: Search Entities Returns Ranked Hits
slug: search-entities-returns-ranked-hits-cbef
relationships:
    - target: search-entities-http-1th3
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:38:08Z"
statement: When a client invokes GET /api/<project>/search, the syded daemon shall respond with 200 OK and a JSON body containing a hits array of matching entities with score and snippet.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: contract:search-entities-http-1th3
requirement_status: active
rationale: The dashboard command palette surfaces ranked search results.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:38:08Z"
---
