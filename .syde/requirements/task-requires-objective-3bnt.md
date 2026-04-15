---
id: REQ-0120
kind: requirement
name: Task Requires Objective
slug: task-requires-objective-3bnt
relationships:
    - target: task-d3oc
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:39Z"
statement: The syde CLI shall require a non-empty objective on every task instance.
req_type: constraint
priority: must
verification: integration test running syde task create without --objective
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Without an objective a task cannot be evaluated for acceptance.
---
