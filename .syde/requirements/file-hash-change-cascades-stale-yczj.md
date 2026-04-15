---
id: REQ-0109
kind: requirement
name: File Hash Change Cascades Stale
slug: file-hash-change-cascades-stale-yczj
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:24Z"
statement: When the hash of a tracked file changes between scans, the syde CLI shall mark that file and all of its ancestor folders stale in the summary tree.
req_type: functional
priority: must
verification: integration test editing a file and asserting ancestor summary_stale flags
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Cascade-up stale marking keeps higher-level summaries honest about drift.
---
