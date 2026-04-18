---
id: REQ-0237
kind: requirement
name: Get Full Tree Returns Scanned At Timestamp
slug: get-full-tree-returns-scanned-at-timestamp-fjs4
relationships:
    - target: get-full-tree-http-c0vt
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:33Z"
statement: When GET /api/<project>/tree succeeds, the syded daemon shall return scanned_at as an ISO8601 string identifying when the tree was scanned.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree
source: manual
source_ref: contract:get-full-tree-http-c0vt
requirement_status: active
rationale: Clients need to detect tree staleness.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:33Z"
---
