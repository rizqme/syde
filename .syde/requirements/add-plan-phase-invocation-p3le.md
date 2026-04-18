---
id: REQ-0255
kind: requirement
name: Add Plan Phase Invocation
slug: add-plan-phase-invocation-p3le
relationships:
    - target: add-plan-phase-fa7g
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:45Z"
statement: When the user runs syde plan add-phase <plan-slug>, the syde CLI shall create a new phase under the named plan and print the allocated phase ID.
req_type: interface
priority: must
verification: integration test invoking syde plan add-phase
source: manual
source_ref: contract:add-plan-phase-fa7g
requirement_status: active
rationale: Plan phases are the unit of grouping for tasks and must be addressable by stable IDs.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:45Z"
---
