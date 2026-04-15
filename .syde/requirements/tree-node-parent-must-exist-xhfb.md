---
id: REQ-0127
kind: requirement
name: Tree Node Parent Must Exist
slug: tree-node-parent-must-exist-xhfb
relationships:
    - target: tree-node-iutv
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:49Z"
statement: If a non-root tree node declares a parent path that is not present in the summary tree, then the syde CLI shall reject the scan with a validation error.
req_type: constraint
priority: must
verification: integration test asserting every non-root node has an existing parent after scan
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Missing parents break ancestor traversal used by cascade-stale marking.
---
