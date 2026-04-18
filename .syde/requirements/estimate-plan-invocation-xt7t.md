---
id: REQ-0269
kind: requirement
name: Estimate Plan Invocation
slug: estimate-plan-invocation-xt7t
relationships:
    - target: estimate-plan-kwt6
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:45:40Z"
statement: When the user runs syde plan estimate <slug>, the syde CLI shall return the task_count and a size recommendation for the named plan.
req_type: interface
priority: must
verification: integration test invoking syde plan estimate
source: manual
source_ref: contract:estimate-plan-kwt6
requirement_status: active
rationale: Size feedback helps operators split overly large plans before execution.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:21Z"
---
