---
id: REQ-0256
kind: requirement
name: Add Plan Phase Requires Plan Slug
slug: add-plan-phase-requires-plan-slug-4ags
relationships:
    - target: add-plan-phase-fa7g
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When syde plan add-phase is invoked, the syde CLI shall accept plan-slug as a required positional string.
req_type: interface
priority: must
verification: integration test invoking syde plan add-phase without a slug
source: manual
source_ref: contract:add-plan-phase-fa7g
requirement_status: active
rationale: A phase must always be scoped to an existing plan.
---
