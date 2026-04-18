---
id: REQ-0290
kind: requirement
name: Query Entity Full Flag
slug: query-entity-full-flag-9h38
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:09Z"
statement: Where --full is passed to syde query, the syde CLI shall include the entity body and all related data in the output.
req_type: interface
priority: must
verification: integration test invoking syde query --full
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: The full flag unlocks the deep-dive view needed for targeted investigation.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:09Z"
---
