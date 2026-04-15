---
id: REQ-0260
kind: requirement
name: Complete Task Invocation
slug: complete-task-invocation-dm0b
relationships:
    - target: complete-task-k8je
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When the user runs syde task done <slug>, the syde CLI shall transition the task's task_status field to completed.
req_type: interface
priority: must
verification: integration test invoking syde task done
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Task completion is the canonical progress signal for plans.
---
