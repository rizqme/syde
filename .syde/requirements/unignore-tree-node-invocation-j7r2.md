---
id: REQ-0315
kind: requirement
name: Unignore Tree Node Invocation
slug: unignore-tree-node-invocation-j7r2
relationships:
    - target: unignore-tree-node-eoyv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:51Z"
statement: When the user runs syde tree unignore <path>, the syde CLI shall remove the ignore flag on the named node and mark it stale for the next summarize pass.
req_type: interface
priority: must
verification: integration test invoking syde tree unignore on a previously ignored path
source: manual
source_ref: contract:unignore-tree-node-eoyv
requirement_status: active
rationale: Unignoring is the reverse operation needed when ignored paths become relevant again.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:51Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:51Z"
---
