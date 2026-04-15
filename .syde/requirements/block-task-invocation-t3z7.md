---
id: REQ-0258
kind: requirement
name: Block Task Invocation
slug: block-task-invocation-t3z7
relationships:
    - target: block-task-egd4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When the user runs syde task block <slug>, the syde CLI shall transition the task's task_status field to blocked.
req_type: interface
priority: must
verification: integration test invoking syde task block and checking task_status
source: manual
source_ref: contract:block-task-egd4
requirement_status: active
rationale: Blocked state communicates that work cannot proceed without unblocking.
---
