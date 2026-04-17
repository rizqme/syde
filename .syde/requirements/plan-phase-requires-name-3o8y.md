---
id: REQ-0057
kind: requirement
name: Plan Phase Requires Name
slug: plan-phase-requires-name-3o8y
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:01Z"
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
---
