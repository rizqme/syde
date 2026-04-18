---
id: REQ-0305
kind: requirement
name: Show Tree Stale Flag
slug: show-tree-stale-flag-o4t2
relationships:
    - target: show-tree-t0as
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:32Z"
statement: Where --stale is passed to syde tree show, the syde CLI shall prefix stale entries with an exclamation mark in the rendered tree.
req_type: interface
priority: must
verification: integration test invoking syde tree show --stale after modifying a file
source: manual
source_ref: contract:show-tree-t0as
requirement_status: active
rationale: Stale markers make it obvious which parts of the tree still need summarizing.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:32Z"
---
