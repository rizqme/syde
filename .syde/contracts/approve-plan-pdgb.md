---
id: CON-0029
kind: contract
name: Approve Plan
slug: approve-plan-pdgb
description: Mark a plan as approved (gated on explicit user chat approval).
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan approve <slug>
input_parameters:
- path: slug
  type: string
  description: positional, required
output: exit 0; transitions plan_status to approved
output_parameters:
- path: new_status
  type: string
  description: approved
---
