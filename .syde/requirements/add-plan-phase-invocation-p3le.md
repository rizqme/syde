---
id: REQ-0255
kind: requirement
name: Add Plan Phase Invocation
slug: add-plan-phase-invocation-p3le
relationships:
    - target: add-plan-phase-fa7g
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When the user runs syde plan add-phase <plan-slug>, the syde CLI shall create a new phase under the named plan and print the allocated phase ID.
req_type: interface
priority: must
verification: integration test invoking syde plan add-phase
source: manual
source_ref: contract:add-plan-phase-fa7g
requirement_status: active
rationale: Plan phases are the unit of grouping for tasks and must be addressable by stable IDs.
---
