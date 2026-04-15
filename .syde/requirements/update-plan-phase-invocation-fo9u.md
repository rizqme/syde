---
id: REQ-0318
kind: requirement
name: Update Plan Phase Invocation
slug: update-plan-phase-invocation-fo9u
relationships:
    - target: update-plan-phase-izh0
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde plan phase <plan-slug> <phase-id>, the syde CLI shall apply the provided flags to the named phase and print the echoed phase id.
req_type: interface
priority: must
verification: integration test invoking syde plan phase with updates
source: manual
source_ref: contract:update-plan-phase-izh0
requirement_status: active
rationale: Phase updates are required when plans evolve during execution.
---
