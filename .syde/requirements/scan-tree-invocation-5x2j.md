---
id: REQ-0296
kind: requirement
name: Scan Tree Invocation
slug: scan-tree-invocation-5x2j
relationships:
    - target: scan-tree-vmkd
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:50Z"
statement: When the user runs syde tree scan, the syde CLI shall update .syde/tree.yaml in place and print counts of added, changed, deleted, and stale nodes.
req_type: interface
priority: must
verification: integration test invoking syde tree scan after editing files
source: manual
source_ref: contract:scan-tree-vmkd
requirement_status: active
rationale: Tree scan is the first step of every sync workflow and must report diffs accurately.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:50Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:50Z"
---
