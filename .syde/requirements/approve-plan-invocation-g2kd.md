---
id: REQ-0257
kind: requirement
name: Approve Plan Invocation
slug: approve-plan-invocation-g2kd
relationships:
    - target: approve-plan-pdgb
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:54Z"
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
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:54Z"
---
