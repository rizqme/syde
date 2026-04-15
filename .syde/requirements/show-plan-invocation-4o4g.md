---
id: REQ-0302
kind: requirement
name: Show Plan Invocation
slug: show-plan-invocation-4o4g
relationships:
    - target: show-plan-1ybx
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde plan show <slug>, the syde CLI shall render the plan as an ASCII tree of phases and tasks.
req_type: interface
priority: must
verification: integration test invoking syde plan show
source: manual
source_ref: contract:show-plan-1ybx
requirement_status: active
rationale: Plan rendering is the primary way operators review plan structure in the terminal.
---
