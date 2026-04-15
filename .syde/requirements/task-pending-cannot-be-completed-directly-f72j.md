---
id: REQ-0122
kind: requirement
name: Task Pending Cannot Be Completed Directly
slug: task-pending-cannot-be-completed-directly-f72j
relationships:
    - target: task-d3oc
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:41Z"
statement: While a task has status pending, the syde CLI shall not allow it to transition directly to completed without first entering in_progress.
req_type: functional
priority: should
verification: integration test running syde task done on a pending task
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Enforces the pending to in_progress to completed ordering so work is tracked.
---
