---
id: REQ-0304
kind: requirement
name: Show Tree Invocation
slug: show-tree-invocation-kkhy
relationships:
    - target: show-tree-t0as
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree show, the syde CLI shall render an ASCII tree with inline summaries on stdout.
req_type: interface
priority: must
verification: integration test invoking syde tree show
source: manual
source_ref: contract:show-tree-t0as
requirement_status: active
rationale: Tree show is the primary summary-level navigation command for the codebase.
---
