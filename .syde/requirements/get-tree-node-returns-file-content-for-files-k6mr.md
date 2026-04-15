---
id: REQ-0248
kind: requirement
name: Get Tree Node Returns File Content For Files
slug: get-tree-node-returns-file-content-for-files-k6mr
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:01Z"
statement: When GET /api/<project>/tree/<path> resolves to a file, the syded daemon shall return content as the raw file content string.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: The file view shows the raw source without a second round trip.
---
