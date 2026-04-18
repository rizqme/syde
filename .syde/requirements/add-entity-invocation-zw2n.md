---
id: REQ-0232
kind: requirement
name: Add Entity Invocation
slug: add-entity-invocation-zw2n
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:55Z"
statement: When the user runs syde add <kind> <name>, the syde CLI shall create a new markdown entity file under .syde/<kind>/ and print the allocated ID and file path.
req_type: interface
priority: must
verification: integration test invoking syde add
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: The add command is the canonical entry point for creating every kind of design entity.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:55Z"
---
