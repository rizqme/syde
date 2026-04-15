---
id: REQ-0191
kind: requirement
name: Summary Tree Clears Stale Bit On Set Summary
slug: summary-tree-clears-stale-bit-on-set-summary-ya4v
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:58Z"
statement: When SetSummary writes a new summary for a node, the summary tree shall clear the stale bit on that node while marking its parent stale.
req_type: functional
priority: must
verification: unit test of SetSummary in internal/tree/summarize.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Leaf-first summarization progresses upward without losing parent tracking.
---
