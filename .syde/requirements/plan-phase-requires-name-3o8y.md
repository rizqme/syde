---
id: REQ-0057
kind: requirement
name: Plan Phase Requires Name
slug: plan-phase-requires-name-3o8y
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:33Z"
statement: The syde CLI shall require a non-empty name on every plan phase instance.
req_type: constraint
priority: must
verification: integration test running syde plan add-phase without --name
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Phases without names cannot be referenced by tasks or in reports.
---
