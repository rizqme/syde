---
id: REQ-0285
kind: requirement
name: List Tree Changes Invocation
slug: list-tree-changes-invocation-toin
relationships:
    - target: list-tree-changes-mqhv
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde tree changes, the syde CLI shall return a sorted list of stale path nodes with the deepest paths listed first.
req_type: interface
priority: must
verification: integration test invoking syde tree changes after modifying files
source: manual
source_ref: contract:list-tree-changes-mqhv
requirement_status: active
rationale: Deepest-first ordering drives the leaves-first summarize workflow.
---
