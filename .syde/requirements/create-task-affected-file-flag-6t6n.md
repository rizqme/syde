---
id: REQ-0268
kind: requirement
name: Create Task Affected File Flag
slug: create-task-affected-file-flag-6t6n
relationships:
    - target: create-task-23f4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When syde task create is invoked, the syde CLI shall accept --affected-file as a repeatable string whose values must exist in the tracked tree.
req_type: interface
priority: must
verification: integration test invoking syde task create --affected-file with a non-tracked path
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Validating affected files prevents tasks from referencing phantom paths.
---
