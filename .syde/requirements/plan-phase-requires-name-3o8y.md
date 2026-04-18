---
id: REQ-0057
kind: requirement
name: Plan Phase Requires Name
slug: plan-phase-requires-name-3o8y
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:41Z"
statement: The syde CLI shall require a non-empty name on every plan phase instance.
req_type: constraint
priority: must
verification: integration test running syde plan add-phase without --name
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Phases without names cannot be referenced by tasks or in reports.
audited_overlaps:
    - slug: entity-requires-description-iys8
      distinction: Requires a non-empty name field on plan phases; target requires a non-empty description on every entity.
    - slug: plan-requires-objective-vacv
      distinction: Requires non-empty name on plan phases; target requires non-empty objective on the plan itself.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:41Z"
---
