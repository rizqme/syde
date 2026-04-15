---
id: REQ-0189
kind: requirement
name: Summary Tree Auto Summarizes Binary And Large Files
slug: summary-tree-auto-summarizes-binary-and-large-files-sdm3
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:50Z"
statement: When a walked file is binary or larger than 1 MiB, the summary tree shall auto-generate a summary so the file does not appear on the stale list.
req_type: functional
priority: must
verification: unit test covering binary and large-file handling in scan.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Auto-summaries prevent humans from having to narrate non-text blobs.
---
