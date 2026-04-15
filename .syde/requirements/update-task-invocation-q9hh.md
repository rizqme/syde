---
id: REQ-0321
kind: requirement
name: Update Task Invocation
slug: update-task-invocation-q9hh
relationships:
    - target: update-task-6y9m
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde task update <slug>, the syde CLI shall apply the provided field flags to the named task and print its updated slug.
req_type: interface
priority: must
verification: integration test invoking syde task update
source: manual
source_ref: contract:update-task-6y9m
requirement_status: active
rationale: Task updates are required as scope, acceptance, and affected entities shift during execution.
---
