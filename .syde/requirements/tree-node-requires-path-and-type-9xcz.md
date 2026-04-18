---
id: REQ-0130
kind: requirement
name: Tree Node Requires Path And Type
slug: tree-node-requires-path-and-type-9xcz
relationships:
    - target: tree-node-iutv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:08Z"
statement: The syde CLI shall require both a path and a type of file or dir on every tree node instance.
req_type: constraint
priority: must
verification: unit test rejecting a tree node missing path or type
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Path and type are the minimum identity fields required to render and traverse the tree.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:08Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:38:08Z"
---
