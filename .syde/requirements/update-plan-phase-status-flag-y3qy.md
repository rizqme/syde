---
id: REQ-0319
kind: requirement
name: Update Plan Phase Status Flag
slug: update-plan-phase-status-flag-y3qy
relationships:
    - target: update-plan-phase-izh0
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When syde plan phase is invoked with --status, the syde CLI shall accept one of pending, in_progress, completed, or skipped as the new phase status.
req_type: interface
priority: must
verification: integration test invoking syde plan phase --status
source: manual
source_ref: contract:update-plan-phase-izh0
requirement_status: active
rationale: Phase status transitions drive plan progress rollups.
---
