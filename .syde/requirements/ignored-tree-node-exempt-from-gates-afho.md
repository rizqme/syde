---
id: REQ-0132
kind: requirement
name: Ignored Tree Node Exempt From Gates
slug: ignored-tree-node-exempt-from-gates-afho
relationships:
    - target: tree-node-iutv
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:57Z"
statement: While a tree node has ignored set to true, the syde CLI shall exclude it from orphan and stale-tree status gates.
req_type: functional
priority: must
verification: integration test marking a node ignored and running syde tree status --strict
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Ignored nodes are intentional exclusions that must not block sync gates.
---
