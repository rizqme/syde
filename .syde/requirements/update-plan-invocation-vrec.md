---
id: REQ-0320
kind: requirement
name: Update Plan Invocation
slug: update-plan-invocation-vrec
relationships:
    - target: update-plan-s5x9
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-17T10:46:07Z"
statement: When the user runs syde plan update <slug>, the syde CLI shall apply the provided field flags to the named plan and print its updated slug.
req_type: interface
priority: must
verification: integration test invoking syde plan update
source: manual
source_ref: contract:update-plan-s5x9
requirement_status: active
rationale: Plan updates are required as scope, objectives, and background evolve.
audited_overlaps:
    - slug: update-entity-invocation-3p7y
      distinction: syde plan update is the plan-specific command; syde update is the generic entity command with different invocation paths.
    - slug: update-task-invocation-q9hh
      distinction: syde plan update mutates plan entities; syde task update mutates task entities, different commands and target kinds.
---
