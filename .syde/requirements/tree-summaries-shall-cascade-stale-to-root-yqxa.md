---
id: REQ-0005
kind: requirement
name: Tree summaries shall cascade stale to root
slug: tree-summaries-shall-cascade-stale-to-root-yqxa
description: File hash changes propagate stale markers up the summary tree.
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:59Z"
statement: When any file hash changes, the summary tree shall mark that file and every ancestor folder stale up to the root, and when any summary is updated the direct parent shall be marked stale.
req_type: functional
priority: must
verification: syde tree scan after a single file edit marks the file plus every ancestor folder stale
source: manual
source_ref: decision:DEC-0005
requirement_status: active
rationale: Folder summaries are derived from their children. If a child changes, the folder's summary is by definition out of date. Cascading stale up ensures no one forgets to re-summarize the parent.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:59Z"
---
