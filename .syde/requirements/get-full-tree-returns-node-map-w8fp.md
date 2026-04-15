---
id: REQ-0235
kind: requirement
name: Get Full Tree Returns Node Map
slug: get-full-tree-returns-node-map-w8fp
relationships:
    - target: get-full-tree-http-c0vt
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:50Z"
statement: When a client invokes GET /api/<project>/tree, the syded daemon shall respond with 200 OK and a JSON body containing a path-keyed nodes map and a scanned_at timestamp.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: contract:get-full-tree-http-c0vt
requirement_status: active
rationale: The FileTree page renders from the flat node map.
---
