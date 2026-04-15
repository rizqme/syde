---
id: REQ-0046
kind: requirement
name: HTTP API Serves Search
slug: http-api-serves-search-xcl1
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: When a client invokes the search endpoint, the syded daemon shall respond with matching entities across all kinds for the given query.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The command palette depends on a server-side search endpoint.
---
