---
id: REQ-0120
kind: requirement
name: Task Requires Objective
slug: task-requires-objective-3bnt
relationships:
    - target: task-d3oc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:48Z"
statement: The syde CLI shall require a non-empty objective on every task instance.
req_type: constraint
priority: must
verification: integration test running syde task create without --objective
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Without an objective a task cannot be evaluated for acceptance.
audited_overlaps:
    - slug: entity-requires-description-iys8
      distinction: Task objective is a task-specific required field, separate from the generic entity description mandated across all entity kinds.
    - slug: plan-requires-objective-vacv
      distinction: Task objective applies to task entities; plan objective applies to plan entities, different entity kinds despite shared field name.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:48Z"
---
