---
id: REQ-0270
kind: requirement
name: Execute Plan Invocation
slug: execute-plan-invocation-f8md
relationships:
    - target: execute-plan-50cg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:14Z"
statement: When the user runs syde plan execute <slug>, the syde CLI shall transition the plan's plan_status field to in-progress.
req_type: interface
priority: must
verification: integration test invoking syde plan execute on an approved plan
source: manual
source_ref: contract:execute-plan-50cg
requirement_status: active
rationale: Execution state signals that implementation work is underway against the plan.
audited_overlaps:
    - slug: approve-plan-invocation-g2kd
      distinction: Transitions plan_status to in-progress via syde plan execute, not to approved via syde plan approve.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:14Z"
---
