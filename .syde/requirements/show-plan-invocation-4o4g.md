---
id: REQ-0302
kind: requirement
name: Show Plan Invocation
slug: show-plan-invocation-4o4g
relationships:
    - target: show-plan-1ybx
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:51Z"
statement: When the user runs syde plan show <slug>, the syde CLI shall render the plan as an ASCII tree of phases and tasks.
req_type: interface
priority: must
verification: integration test invoking syde plan show
source: manual
source_ref: contract:show-plan-1ybx
requirement_status: active
rationale: Plan rendering is the primary way operators review plan structure in the terminal.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:51Z"
---
