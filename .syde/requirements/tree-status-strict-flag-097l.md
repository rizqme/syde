---
id: REQ-0314
kind: requirement
name: Tree Status Strict Flag
slug: tree-status-strict-flag-097l
relationships:
    - target: tree-status-k6ag
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: Where --strict is passed to syde tree status, the syde CLI shall exit with code 1 whenever any tracked node is stale.
req_type: interface
priority: must
verification: integration test invoking syde tree status --strict with stale nodes present
source: manual
source_ref: contract:tree-status-k6ag
requirement_status: active
rationale: Strict mode enables CI and session-end hooks to block on dirty trees.
---
