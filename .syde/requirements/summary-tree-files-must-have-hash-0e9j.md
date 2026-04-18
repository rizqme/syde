---
id: REQ-0108
kind: requirement
name: Summary Tree Files Must Have Hash
slug: summary-tree-files-must-have-hash-0e9j
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:43Z"
statement: The syde CLI shall ensure that every tracked file node in the summary tree carries a SHA-256 content hash.
req_type: constraint
priority: must
verification: integration test running syde tree scan and asserting every file node has a hash
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: Hashes are the basis for stale detection and cascade-up marking.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:43Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:43Z"
---
