---
id: REQ-0257
kind: requirement
name: Approve Plan Invocation
slug: approve-plan-invocation-g2kd
relationships:
    - target: approve-plan-pdgb
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-17T10:45:22Z"
statement: When the user runs syde plan approve <slug>, the syde CLI shall transition the plan's plan_status field to approved.
req_type: interface
priority: must
verification: integration test invoking syde plan approve and checking plan_status
source: manual
source_ref: contract:approve-plan-pdgb
requirement_status: active
rationale: Approval gates plan execution and downstream task work.
audited_overlaps:
    - slug: execute-plan-invocation-f8md
      distinction: Different CLI subcommand (approve vs execute) driving different plan_status transitions (approved vs in-progress) at distinct workflow stages.
---
