---
id: REQ-0291
kind: requirement
name: Query Entity Impacts Flag
slug: query-entity-impacts-flag-f98d
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: graph-engine-xgjy
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:45Z"
statement: Where --impacts is passed to syde query, the syde CLI shall return a transitive impact analysis starting at the named slug.
req_type: interface
priority: must
verification: integration test invoking syde query --impacts on a referenced entity
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: Impact analysis supports change-risk reviews before editing architecturally central entities.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:45Z"
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:37:45Z"
---
