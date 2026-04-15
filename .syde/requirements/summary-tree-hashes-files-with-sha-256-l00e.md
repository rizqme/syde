---
id: REQ-0186
kind: requirement
name: Summary Tree Hashes Files With SHA 256
slug: summary-tree-hashes-files-with-sha-256-l00e
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:41Z"
statement: When the summary tree walks the project, the summary tree shall compute a SHA-256 hash for every non-ignored file.
req_type: functional
priority: must
verification: unit test of WalkProject in internal/tree/walk.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Content hashing is the basis for change detection.
---
