---
id: REQ-0232
kind: requirement
name: Add Entity Invocation
slug: add-entity-invocation-zw2n
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T10:59:47Z"
statement: When the user runs syde add <kind> <name>, the syde CLI shall create a new markdown entity file under .syde/<kind>/ and print the allocated ID and file path.
req_type: interface
priority: must
verification: integration test invoking syde add
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: The add command is the canonical entry point for creating every kind of design entity.
---
