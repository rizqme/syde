---
id: REQ-0319
kind: requirement
name: Update Plan Phase Status Flag
slug: update-plan-phase-status-flag-y3qy
relationships:
    - target: update-plan-phase-izh0
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:57Z"
statement: When syde plan phase is invoked with --status, the syde CLI shall accept one of pending, in_progress, completed, or skipped as the new phase status.
req_type: interface
priority: must
verification: integration test invoking syde plan phase --status
source: manual
source_ref: contract:update-plan-phase-izh0
requirement_status: active
rationale: Phase status transitions drive plan progress rollups.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:57Z"
---
