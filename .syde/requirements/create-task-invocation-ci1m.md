---
id: REQ-0266
kind: requirement
name: Create Task Invocation
slug: create-task-invocation-ci1m
relationships:
    - target: create-task-23f4
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:40Z"
statement: When the user runs syde task create <name>, the syde CLI shall create a new task entity file and print its slug.
req_type: interface
priority: must
verification: integration test invoking syde task create
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Task creation is the primary mechanism for tracking implementation work.
audited_overlaps:
    - slug: create-plan-invocation-inv4
      distinction: Covers syde task create, producing a task entity and printing its slug; the plan creation requirement covers a different subcommand producing a plan entity and also prints a plan ID.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:40Z"
---
