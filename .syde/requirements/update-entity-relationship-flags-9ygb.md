---
id: REQ-0317
kind: requirement
name: Update Entity Relationship Flags
slug: update-entity-relationship-flags-9ygb
relationships:
    - target: update-entity-zpnh
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:45Z"
statement: When syde update is invoked, the syde CLI shall accept --add-rel as a repeatable target-slug-and-type pair and --remove-rel as a repeatable target slug.
req_type: interface
priority: must
verification: integration test invoking syde update --add-rel and --remove-rel
source: manual
source_ref: contract:update-entity-zpnh
requirement_status: active
rationale: Adjusting relationships is how operators reshape the entity graph without rewriting markdown by hand.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:45Z"
---
