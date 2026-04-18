---
id: CON-0025
kind: contract
name: Create Plan
slug: create-plan-t3mn
description: Create a new draft plan with background, objective, and scope.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan create <name> [flags]
input_parameters:
- path: name
  type: string
  description: positional, required. Plan name
- path: --background
  type: string
  description: why this plan exists
- path: --objective
  type: string
  description: what success looks like
- path: --scope
  type: string
  description: in-scope and out-of-scope summary
- path: --description
  type: string
  description: short description
- path: --purpose
  type: string
  description: why it exists
output: exit 0; prints plan slug and file path
output_parameters:
- path: plan_id
  type: string
  description: generated plan ID (PLN-NNNN)
- path: slug
  type: string
  description: file-level slug with -XXXX suffix
---
