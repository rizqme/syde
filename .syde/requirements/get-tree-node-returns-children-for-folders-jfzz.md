---
id: REQ-0250
kind: requirement
name: Get Tree Node Returns Children For Folders
slug: get-tree-node-returns-children-for-folders-jfzz
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:03Z"
statement: When GET /api/<project>/tree/<path> resolves to a folder, the syded daemon shall return children as an array of direct child node objects.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: Folder views expand to show direct children.
---
