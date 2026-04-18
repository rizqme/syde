---
id: REQ-0235
kind: requirement
name: Get Full Tree Returns Node Map
slug: get-full-tree-returns-node-map-w8fp
relationships:
    - target: get-full-tree-http-c0vt
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:15Z"
statement: When a client invokes GET /api/<project>/tree, the syded daemon shall respond with 200 OK and a JSON body containing a path-keyed nodes map and a scanned_at timestamp.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: contract:get-full-tree-http-c0vt
requirement_status: active
rationale: The FileTree page renders from the flat node map.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:15Z"
---
