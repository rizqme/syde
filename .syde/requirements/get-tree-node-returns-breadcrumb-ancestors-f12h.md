---
id: REQ-0246
kind: requirement
name: Get Tree Node Returns Breadcrumb Ancestors
slug: get-tree-node-returns-breadcrumb-ancestors-f12h
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:59Z"
statement: When GET /api/<project>/tree/<path> succeeds, the syded daemon shall return breadcrumb as an array of ancestor folder summary objects.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: The tree node view renders an ancestor path breadcrumb.
---
