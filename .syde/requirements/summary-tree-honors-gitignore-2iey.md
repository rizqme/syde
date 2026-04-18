---
id: REQ-0188
kind: requirement
name: Summary Tree Honors gitignore
slug: summary-tree-honors-gitignore-2iey
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:21Z"
statement: When walking the project, the summary tree shall honor .gitignore and the tree_ignore patterns in syde.yaml in addition to built-in ignore rules.
req_type: functional
priority: must
verification: unit test of ignore matching in internal/tree/ignore.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Ignoring generated and vendored files keeps the tree focused on source.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:21Z"
---
