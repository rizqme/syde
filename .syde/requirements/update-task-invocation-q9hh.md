---
id: REQ-0321
kind: requirement
name: Update Task Invocation
slug: update-task-invocation-q9hh
relationships:
    - target: update-task-6y9m
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:06Z"
statement: When the user runs syde task update <slug>, the syde CLI shall apply the provided field flags to the named task and print its updated slug.
req_type: interface
priority: must
verification: integration test invoking syde task update
source: manual
source_ref: contract:update-task-6y9m
requirement_status: active
rationale: Task updates are required as scope, acceptance, and affected entities shift during execution.
audited_overlaps:
    - slug: update-entity-invocation-3p7y
      distinction: syde task update is the task-specific command; syde update is the generic entity command, different invocation paths.
    - slug: update-plan-invocation-vrec
      distinction: syde task update mutates task entities; syde plan update mutates plan entities, different commands and target kinds.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:06Z"
---
