---
id: REQ-0074
kind: requirement
name: Plan Requires Objective
slug: plan-requires-objective-vacv
relationships:
    - target: plan-sk33
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:00Z"
statement: The syde CLI shall require a non-empty objective on every plan instance.
req_type: constraint
priority: must
verification: integration test running syde plan create without --objective
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: A plan without an objective cannot be evaluated for completion.
audited_overlaps:
    - slug: entity-requires-description-iys8
      distinction: Plan objective is a plan-scoped required field distinct from the generic entity description mandate applied to all kinds.
    - slug: plan-phase-requires-name-3o8y
      distinction: Plan-level objective differs from the phase-level name field required on each plan phase sub-instance.
    - slug: task-requires-objective-3bnt
      distinction: Plan objective governs plan entities; task objective is a separate required field on task entities.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:00Z"
---
