---
contract_kind: cli
description: Mark a plan as approved (gated on explicit user chat approval).
id: CON-0029
input: syde plan approve <slug>
input_parameters:
    - description: positional, required
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Approve Plan
output: exit 0; transitions plan_status to approved
output_parameters:
    - description: approved
      path: new_status
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: approve-plan-pdgb
updated_at: "2026-04-14T03:27:04Z"
---
