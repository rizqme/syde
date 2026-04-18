---
id: REQ-0244
kind: requirement
name: Add Entity Name Arg
slug: add-entity-name-arg-0siy
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:31Z"
statement: When syde add is invoked, the syde CLI shall require a positional name argument used to derive the entity slug.
req_type: interface
priority: must
verification: integration test invoking syde add with missing name
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: The entity name drives slug generation and human-readable identification.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:31Z"
---
