---
id: REQ-0271
kind: requirement
name: Execute Plan Requires Approved Status
slug: execute-plan-requires-approved-status-ced1
relationships:
    - target: execute-plan-50cg
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: If syde plan execute is invoked on a plan that is not in approved status, then the syde CLI shall reject the command and leave the plan_status unchanged.
req_type: interface
priority: must
verification: integration test invoking syde plan execute on an unapproved plan
source: manual
source_ref: contract:execute-plan-50cg
requirement_status: active
rationale: Approval gating prevents execution of plans that have not been reviewed.
---
