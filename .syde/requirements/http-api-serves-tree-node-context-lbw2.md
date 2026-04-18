---
id: REQ-0045
kind: requirement
name: HTTP API Serves Tree Node Context
slug: http-api-serves-tree-node-context-lbw2
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:48Z"
statement: When a client invokes GET /api/<project>/tree/<path>, the syded daemon shall respond with the context bundle for that node.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Drill-down views need per-path context bundles without refetching the entire tree.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:48Z"
---
