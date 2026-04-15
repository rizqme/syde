---
id: REQ-0005
kind: requirement
name: Tree summaries shall cascade stale to root
slug: tree-summaries-shall-cascade-stale-to-root-yqxa
description: File hash changes propagate stale markers up the summary tree.
relationships:
    - target: summary-tree-4ksz
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:47:02Z"
statement: When any file hash changes, the summary tree shall mark that file and every ancestor folder stale up to the root, and when any summary is updated the direct parent shall be marked stale.
req_type: functional
priority: must
verification: syde tree scan after a single file edit marks the file plus every ancestor folder stale
source: manual
source_ref: decision:DEC-0005
requirement_status: active
rationale: Folder summaries are derived from their children. If a child changes, the folder's summary is by definition out of date. Cascading stale up ensures no one forgets to re-summarize the parent.
---
