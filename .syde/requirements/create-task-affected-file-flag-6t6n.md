---
id: REQ-0268
kind: requirement
name: Create Task Affected File Flag
slug: create-task-affected-file-flag-6t6n
relationships:
    - target: create-task-23f4
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:11Z"
statement: When syde task create is invoked, the syde CLI shall accept --affected-file as a repeatable string whose values must exist in the tracked tree.
req_type: interface
priority: must
verification: integration test invoking syde task create --affected-file with a non-tracked path
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Validating affected files prevents tasks from referencing phantom paths.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:11Z"
---
