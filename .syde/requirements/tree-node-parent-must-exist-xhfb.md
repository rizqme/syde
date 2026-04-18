---
id: REQ-0127
kind: requirement
name: Tree Node Parent Must Exist
slug: tree-node-parent-must-exist-xhfb
relationships:
    - target: tree-node-iutv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:04Z"
statement: If a non-root tree node declares a parent path that is not present in the summary tree, then the syde CLI shall reject the scan with a validation error.
req_type: constraint
priority: must
verification: integration test asserting every non-root node has an existing parent after scan
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Missing parents break ancestor traversal used by cascade-stale marking.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:04Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:38:04Z"
---
