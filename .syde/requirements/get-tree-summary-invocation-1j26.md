---
id: REQ-0274
kind: requirement
name: Get Tree Summary Invocation
slug: get-tree-summary-invocation-1j26
relationships:
    - target: get-tree-summary-2vyd
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:10Z"
statement: When the user runs syde tree get <path>, the syde CLI shall return the stored summary text for the named tree node.
req_type: interface
priority: must
verification: integration test invoking syde tree get on a summarized path
source: manual
source_ref: contract:get-tree-summary-2vyd
requirement_status: active
rationale: Targeted summary retrieval lets agents inspect single nodes without rendering the full tree.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:10Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:10Z"
---
