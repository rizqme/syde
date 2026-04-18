---
id: REQ-0299
kind: requirement
name: Set Tree Summary Invocation
slug: set-tree-summary-invocation-xcw0
relationships:
    - target: set-tree-summary-2vdt
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:09Z"
statement: When the user runs syde tree summarize <path> --summary <text>, the syde CLI shall store the summary on the named node and mark its parent stale.
req_type: interface
priority: must
verification: integration test invoking syde tree summarize and checking parent stale state
source: manual
source_ref: contract:set-tree-summary-2vdt
requirement_status: active
rationale: Upward stale cascading drives the leaves-first summarize workflow.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:09Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:09Z"
---
