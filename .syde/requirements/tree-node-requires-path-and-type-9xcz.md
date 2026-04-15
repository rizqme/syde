---
id: REQ-0130
kind: requirement
name: Tree Node Requires Path And Type
slug: tree-node-requires-path-and-type-9xcz
relationships:
    - target: tree-node-iutv
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:54Z"
statement: The syde CLI shall require both a path and a type of file or dir on every tree node instance.
req_type: constraint
priority: must
verification: unit test rejecting a tree node missing path or type
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Path and type are the minimum identity fields required to render and traverse the tree.
---
