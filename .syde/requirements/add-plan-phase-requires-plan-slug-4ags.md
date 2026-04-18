---
id: REQ-0256
kind: requirement
name: Add Plan Phase Requires Plan Slug
slug: add-plan-phase-requires-plan-slug-4ags
relationships:
    - target: add-plan-phase-fa7g
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:07Z"
statement: When syde plan add-phase is invoked, the syde CLI shall accept plan-slug as a required positional string.
req_type: interface
priority: must
verification: integration test invoking syde plan add-phase without a slug
source: manual
source_ref: contract:add-plan-phase-fa7g
requirement_status: active
rationale: A phase must always be scoped to an existing plan.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:07Z"
---
