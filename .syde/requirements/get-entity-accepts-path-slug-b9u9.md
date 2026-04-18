---
id: REQ-0273
kind: requirement
name: Get Entity Accepts Path Slug
slug: get-entity-accepts-path-slug-b9u9
relationships:
    - target: get-entity-0lzq
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:16Z"
statement: When syde get is invoked, the syde CLI shall accept slug as a full slug, bare slug, or parent/child path string.
req_type: interface
priority: must
verification: integration test invoking syde get with each slug form
source: manual
source_ref: contract:get-entity-0lzq
requirement_status: active
rationale: Flexible slug resolution reduces friction when operators remember only part of an identifier.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:16Z"
---
