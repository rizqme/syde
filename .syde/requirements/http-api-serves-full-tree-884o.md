---
id: REQ-0044
kind: requirement
name: HTTP API Serves Full Tree
slug: http-api-serves-full-tree-884o
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:25Z"
statement: When a client invokes GET /api/<project>/tree, the syded daemon shall respond with the project's summary tree as JSON.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The file tree screen needs the full tree payload on mount.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:25Z"
---
