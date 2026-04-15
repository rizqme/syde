---
id: REQ-0265
kind: requirement
name: Create Subtask Invocation
slug: create-subtask-invocation-wcw2
relationships:
    - target: create-subtask-0fj9
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde task sub <parent-slug> <name>, the syde CLI shall create a new subtask nested under the named parent task and print its slug.
req_type: interface
priority: must
verification: integration test invoking syde task sub
source: manual
source_ref: contract:create-subtask-0fj9
requirement_status: active
rationale: Subtasks allow decomposition of larger tasks while preserving parent-child traceability.
---
