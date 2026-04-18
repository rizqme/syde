---
id: REQ-0303
kind: requirement
name: Show Plan Full Flag
slug: show-plan-full-flag-foqa
relationships:
    - target: show-plan-1ybx
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:55Z"
statement: Where --full is passed to syde plan show, the syde CLI shall include per-phase details and per-task objective, details, and acceptance fields in the output.
req_type: interface
priority: must
verification: integration test invoking syde plan show --full
source: manual
source_ref: contract:show-plan-1ybx
requirement_status: active
rationale: Full mode exposes the details needed for in-depth plan review.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:55Z"
---
