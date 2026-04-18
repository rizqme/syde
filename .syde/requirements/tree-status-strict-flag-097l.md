---
id: REQ-0314
kind: requirement
name: Tree Status Strict Flag
slug: tree-status-strict-flag-097l
relationships:
    - target: tree-status-k6ag
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:51Z"
statement: Where --strict is passed to syde tree status, the syde CLI shall exit with code 1 whenever any tracked node is stale.
req_type: interface
priority: must
verification: integration test invoking syde tree status --strict with stale nodes present
source: manual
source_ref: contract:tree-status-k6ag
requirement_status: active
rationale: Strict mode enables CI and session-end hooks to block on dirty trees.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:51Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:51Z"
---
