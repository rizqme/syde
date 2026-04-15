---
id: REQ-0317
kind: requirement
name: Update Entity Relationship Flags
slug: update-entity-relationship-flags-9ygb
relationships:
    - target: update-entity-zpnh
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When syde update is invoked, the syde CLI shall accept --add-rel as a repeatable target-slug-and-type pair and --remove-rel as a repeatable target slug.
req_type: interface
priority: must
verification: integration test invoking syde update --add-rel and --remove-rel
source: manual
source_ref: contract:update-entity-zpnh
requirement_status: active
rationale: Adjusting relationships is how operators reshape the entity graph without rewriting markdown by hand.
---
