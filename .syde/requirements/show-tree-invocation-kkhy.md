---
id: REQ-0304
kind: requirement
name: Show Tree Invocation
slug: show-tree-invocation-kkhy
relationships:
    - target: show-tree-t0as
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:20Z"
statement: When the user runs syde tree show, the syde CLI shall render an ASCII tree with inline summaries on stdout.
req_type: interface
priority: must
verification: integration test invoking syde tree show
source: manual
source_ref: contract:show-tree-t0as
requirement_status: active
rationale: Tree show is the primary summary-level navigation command for the codebase.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:20Z"
---
