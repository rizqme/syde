---
id: REQ-0269
kind: requirement
name: Estimate Plan Invocation
slug: estimate-plan-invocation-xt7t
relationships:
    - target: estimate-plan-kwt6
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde plan estimate <slug>, the syde CLI shall return the task_count and a size recommendation for the named plan.
req_type: interface
priority: must
verification: integration test invoking syde plan estimate
source: manual
source_ref: contract:estimate-plan-kwt6
requirement_status: active
rationale: Size feedback helps operators split overly large plans before execution.
---
