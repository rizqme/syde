---
id: REQ-0283
kind: requirement
name: List Tasks Invocation
slug: list-tasks-invocation-ckfj
relationships:
    - target: list-tasks-0sgj
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:43Z"
statement: When the user runs syde task list, the syde CLI shall print a tabular task listing with slug, status, priority, and parent plan.
req_type: interface
priority: must
verification: integration test invoking syde task list
source: manual
source_ref: contract:list-tasks-0sgj
requirement_status: active
rationale: Task listings are the primary view for operators choosing what to work on next.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:43Z"
---
