---
id: REQ-0283
kind: requirement
name: List Tasks Invocation
slug: list-tasks-invocation-ckfj
relationships:
    - target: list-tasks-0sgj
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde task list, the syde CLI shall print a tabular task listing with slug, status, priority, and parent plan.
req_type: interface
priority: must
verification: integration test invoking syde task list
source: manual
source_ref: contract:list-tasks-0sgj
requirement_status: active
rationale: Task listings are the primary view for operators choosing what to work on next.
---
