---
id: REQ-0299
kind: requirement
name: Set Tree Summary Invocation
slug: set-tree-summary-invocation-xcw0
relationships:
    - target: set-tree-summary-2vdt
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree summarize <path> --summary <text>, the syde CLI shall store the summary on the named node and mark its parent stale.
req_type: interface
priority: must
verification: integration test invoking syde tree summarize and checking parent stale state
source: manual
source_ref: contract:set-tree-summary-2vdt
requirement_status: active
rationale: Upward stale cascading drives the leaves-first summarize workflow.
---
