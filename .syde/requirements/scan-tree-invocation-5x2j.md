---
id: REQ-0296
kind: requirement
name: Scan Tree Invocation
slug: scan-tree-invocation-5x2j
relationships:
    - target: scan-tree-vmkd
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree scan, the syde CLI shall update .syde/tree.yaml in place and print counts of added, changed, deleted, and stale nodes.
req_type: interface
priority: must
verification: integration test invoking syde tree scan after editing files
source: manual
source_ref: contract:scan-tree-vmkd
requirement_status: active
rationale: Tree scan is the first step of every sync workflow and must report diffs accurately.
---
