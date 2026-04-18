---
id: REQ-0046
kind: requirement
name: HTTP API Serves Search
slug: http-api-serves-search-xcl1
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:38:08Z"
statement: When a client invokes the search endpoint, the syded daemon shall respond with matching entities across all kinds for the given query.
req_type: interface
priority: must
verification: integration test against /api/<project>/search
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The command palette depends on a server-side search endpoint.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:38:08Z"
---
