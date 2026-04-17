---
id: REQ-0264
kind: requirement
name: Create Plan Invocation
slug: create-plan-invocation-inv4
relationships:
    - target: create-plan-t3mn
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-17T10:46:03Z"
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
---
