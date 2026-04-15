---
id: REQ-0230
kind: requirement
name: Search Entities Returns Ranked Hits
slug: search-entities-returns-ranked-hits-cbef
relationships:
    - target: search-entities-http-1th3
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:43Z"
statement: When a client invokes GET /api/<project>/search, the syded daemon shall respond with 200 OK and a JSON body containing a hits array of matching entities with score and snippet.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: contract:search-entities-http-1th3
requirement_status: active
rationale: The dashboard command palette surfaces ranked search results.
---
