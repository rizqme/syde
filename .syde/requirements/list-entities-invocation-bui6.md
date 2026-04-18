---
id: REQ-0280
kind: requirement
name: List Entities Invocation
slug: list-entities-invocation-bui6
relationships:
    - target: list-entities-0iec
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:57Z"
statement: When the user runs syde list, the syde CLI shall print a tabular listing with one row per entity containing kind, name, and slug.
req_type: interface
priority: must
verification: integration test invoking syde list
source: manual
source_ref: contract:list-entities-0iec
requirement_status: active
rationale: Listing is the fastest way for operators to discover what exists in the model.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:57Z"
---
