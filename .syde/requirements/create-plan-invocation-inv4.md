---
id: REQ-0264
kind: requirement
name: Create Plan Invocation
slug: create-plan-invocation-inv4
relationships:
    - target: create-plan-t3mn
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:59Z"
statement: When the user runs syde plan create <name>, the syde CLI shall create a new plan entity file and print its generated plan ID and slug.
req_type: interface
priority: must
verification: integration test invoking syde plan create
source: manual
source_ref: contract:create-plan-t3mn
requirement_status: active
rationale: Plan creation is the entry point for all design-before-code workflows.
audited_overlaps:
    - slug: create-task-invocation-ci1m
      distinction: Covers syde plan create, producing a plan entity with plan ID and slug; the task creation requirement covers a different subcommand producing a task entity and outputs only a slug.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:59Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:59Z"
---
