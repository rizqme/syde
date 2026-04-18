---
id: REQ-0281
kind: requirement
name: List Entities Kind Filter
slug: list-entities-kind-filter-6c0r
relationships:
    - target: list-entities-0iec
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:24Z"
statement: When syde list is invoked with a positional kind argument, the syde CLI shall filter the listing to entities of that kind.
req_type: interface
priority: must
verification: integration test invoking syde list component
source: manual
source_ref: contract:list-entities-0iec
requirement_status: active
rationale: Kind filtering narrows results in projects with many entities.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:24Z"
---
