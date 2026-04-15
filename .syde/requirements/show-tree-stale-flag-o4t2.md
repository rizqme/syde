---
id: REQ-0305
kind: requirement
name: Show Tree Stale Flag
slug: show-tree-stale-flag-o4t2
relationships:
    - target: show-tree-t0as
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: Where --stale is passed to syde tree show, the syde CLI shall prefix stale entries with an exclamation mark in the rendered tree.
req_type: interface
priority: must
verification: integration test invoking syde tree show --stale after modifying a file
source: manual
source_ref: contract:show-tree-t0as
requirement_status: active
rationale: Stale markers make it obvious which parts of the tree still need summarizing.
---
