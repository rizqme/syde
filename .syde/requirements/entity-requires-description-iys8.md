---
id: REQ-0031
kind: requirement
name: Entity Requires Description
slug: entity-requires-description-iys8
relationships:
    - target: entity-8x6p
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:04Z"
statement: The syde CLI shall require a non-empty description on every entity instance.
req_type: constraint
priority: must
verification: integration test running syde add <kind> without --description
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: The description is the one-sentence elevator pitch used by list and search commands.
audited_overlaps:
    - slug: plan-phase-requires-name-3o8y
      distinction: Applies to the description field on every entity kind, not the name field on plan phase instances specifically.
    - slug: plan-requires-objective-vacv
      distinction: Applies to the description field on every entity kind, not the objective field scoped to plan instances.
    - slug: task-requires-objective-3bnt
      distinction: Applies to the description field on every entity kind, not the objective field scoped to task instances.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:04Z"
---
