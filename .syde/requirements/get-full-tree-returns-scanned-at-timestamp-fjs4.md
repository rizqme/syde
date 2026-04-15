---
id: REQ-0237
kind: requirement
name: Get Full Tree Returns Scanned At Timestamp
slug: get-full-tree-returns-scanned-at-timestamp-fjs4
relationships:
    - target: get-full-tree-http-c0vt
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:52Z"
statement: When GET /api/<project>/tree succeeds, the syded daemon shall return scanned_at as an ISO8601 string identifying when the tree was scanned.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: contract:get-full-tree-http-c0vt
requirement_status: active
rationale: Clients need to detect tree staleness.
---
