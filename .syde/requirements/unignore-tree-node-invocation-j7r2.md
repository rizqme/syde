---
id: REQ-0315
kind: requirement
name: Unignore Tree Node Invocation
slug: unignore-tree-node-invocation-j7r2
relationships:
    - target: unignore-tree-node-eoyv
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree unignore <path>, the syde CLI shall remove the ignore flag on the named node and mark it stale for the next summarize pass.
req_type: interface
priority: must
verification: integration test invoking syde tree unignore on a previously ignored path
source: manual
source_ref: contract:unignore-tree-node-eoyv
requirement_status: active
rationale: Unignoring is the reverse operation needed when ignored paths become relevant again.
---
