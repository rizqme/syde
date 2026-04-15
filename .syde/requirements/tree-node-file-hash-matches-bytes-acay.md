---
id: REQ-0128
kind: requirement
name: Tree Node File Hash Matches Bytes
slug: tree-node-file-hash-matches-bytes-acay
relationships:
    - target: tree-node-iutv
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:52Z"
statement: When syde tree scan runs, the syde CLI shall ensure every file tree node's hash matches the SHA-256 of its current bytes.
req_type: constraint
priority: must
verification: integration test running scan twice with unchanged files and asserting stable hashes
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Hash drift from file bytes would silently corrupt stale detection.
---
