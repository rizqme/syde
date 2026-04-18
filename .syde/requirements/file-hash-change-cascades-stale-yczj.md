---
id: REQ-0109
kind: requirement
name: File Hash Change Cascades Stale
slug: file-hash-change-cascades-stale-yczj
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:46Z"
statement: When the hash of a tracked file changes between scans, the syde CLI shall mark that file and all of its ancestor folders stale in the summary tree.
req_type: functional
priority: must
verification: integration test editing a file and asserting ancestor summary_stale flags
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Cascade-up stale marking keeps higher-level summaries honest about drift.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:46Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:46Z"
---
