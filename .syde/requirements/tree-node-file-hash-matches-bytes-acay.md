---
id: REQ-0128
kind: requirement
name: Tree Node File Hash Matches Bytes
slug: tree-node-file-hash-matches-bytes-acay
relationships:
    - target: tree-node-iutv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:02Z"
statement: When syde tree scan runs, the syde CLI shall ensure every file tree node's hash matches the SHA-256 of its current bytes.
req_type: constraint
priority: must
verification: integration test running scan twice with unchanged files and asserting stable hashes
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Hash drift from file bytes would silently corrupt stale detection.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:02Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:38:02Z"
---
