---
id: REQ-0111
kind: requirement
name: Folder Summary Stale On Child Change
slug: folder-summary-stale-on-child-change-o3yk
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:50Z"
statement: When any child node's summary changes, the syde CLI shall mark its parent folder's summary as stale in the summary tree.
req_type: functional
priority: must
verification: integration test editing a child summary and asserting parent summary_stale
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Folder summaries are derived from children so they must invalidate together.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:50Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:50Z"
---
