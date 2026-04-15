---
id: REQ-0244
kind: requirement
name: Add Entity Name Arg
slug: add-entity-name-arg-0siy
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T10:59:59Z"
statement: When syde add is invoked, the syde CLI shall require a positional name argument used to derive the entity slug.
req_type: interface
priority: must
verification: integration test invoking syde add with missing name
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: The entity name drives slug generation and human-readable identification.
---
