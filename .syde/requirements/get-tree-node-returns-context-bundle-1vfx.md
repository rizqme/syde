---
id: REQ-0240
kind: requirement
name: Get Tree Node Returns Context Bundle
slug: get-tree-node-returns-context-bundle-1vfx
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:54Z"
statement: When a client invokes GET /api/<project>/tree/<path>, the syded daemon shall respond with 200 OK and a JSON ContextBundle containing breadcrumb, summary, content, and children fields.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: Tree node drill-down needs a single bundled response.
---
