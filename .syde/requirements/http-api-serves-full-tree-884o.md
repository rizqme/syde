---
id: REQ-0044
kind: requirement
name: HTTP API Serves Full Tree
slug: http-api-serves-full-tree-884o
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: When a client invokes GET /api/<project>/tree, the syded daemon shall respond with the project's summary tree as JSON.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The file tree screen needs the full tree payload on mount.
---
