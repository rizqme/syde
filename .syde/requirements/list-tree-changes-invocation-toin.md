---
id: REQ-0285
kind: requirement
name: List Tree Changes Invocation
slug: list-tree-changes-invocation-toin
relationships:
    - target: list-tree-changes-mqhv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:58Z"
statement: When the user runs syde tree changes, the syde CLI shall return a sorted list of stale path nodes with the deepest paths listed first.
req_type: interface
priority: must
verification: integration test invoking syde tree changes after modifying files
source: manual
source_ref: contract:list-tree-changes-mqhv
requirement_status: active
rationale: Deepest-first ordering drives the leaves-first summarize workflow.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:58Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:58Z"
---
