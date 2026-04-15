---
id: REQ-0045
kind: requirement
name: HTTP API Serves Tree Node Context
slug: http-api-serves-tree-node-context-lbw2
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: When a client invokes GET /api/<project>/tree/<path>, the syded daemon shall respond with the context bundle for that node.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Drill-down views need per-path context bundles without refetching the entire tree.
---
