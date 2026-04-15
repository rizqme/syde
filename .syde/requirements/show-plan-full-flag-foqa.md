---
id: REQ-0303
kind: requirement
name: Show Plan Full Flag
slug: show-plan-full-flag-foqa
relationships:
    - target: show-plan-1ybx
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: Where --full is passed to syde plan show, the syde CLI shall include per-phase details and per-task objective, details, and acceptance fields in the output.
req_type: interface
priority: must
verification: integration test invoking syde plan show --full
source: manual
source_ref: contract:show-plan-1ybx
requirement_status: active
rationale: Full mode exposes the details needed for in-depth plan review.
---
