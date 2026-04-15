---
id: REQ-0108
kind: requirement
name: Summary Tree Files Must Have Hash
slug: summary-tree-files-must-have-hash-0e9j
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:20Z"
statement: The syde CLI shall ensure that every tracked file node in the summary tree carries a SHA-256 content hash.
req_type: constraint
priority: must
verification: integration test running syde tree scan and asserting every file node has a hash
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Hashes are the basis for stale detection and cascade-up marking.
---
