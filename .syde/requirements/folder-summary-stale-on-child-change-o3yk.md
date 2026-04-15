---
id: REQ-0111
kind: requirement
name: Folder Summary Stale On Child Change
slug: folder-summary-stale-on-child-change-o3yk
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:25Z"
statement: When any child node's summary changes, the syde CLI shall mark its parent folder's summary as stale in the summary tree.
req_type: functional
priority: must
verification: integration test editing a child summary and asserting parent summary_stale
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Folder summaries are derived from children so they must invalidate together.
---
