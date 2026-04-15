---
id: REQ-0266
kind: requirement
name: Create Task Invocation
slug: create-task-invocation-ci1m
relationships:
    - target: create-task-23f4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde task create <name>, the syde CLI shall create a new task entity file and print its slug.
req_type: interface
priority: must
verification: integration test invoking syde task create
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Task creation is the primary mechanism for tracking implementation work.
---
