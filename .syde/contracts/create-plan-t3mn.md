---
contract_kind: cli
description: Create a new draft plan with background, objective, and scope.
id: CON-0025
input: syde plan create <name> [flags]
input_parameters:
    - description: positional, required. Plan name
      path: name
      type: string
    - description: why this plan exists
      path: --background
      type: string
    - description: what success looks like
      path: --objective
      type: string
    - description: in-scope and out-of-scope summary
      path: --scope
      type: string
    - description: short description
      path: --description
      type: string
    - description: why it exists
      path: --purpose
      type: string
interaction_pattern: request-response
kind: contract
name: Create Plan
output: exit 0; prints plan slug and file path
output_parameters:
    - description: generated plan ID (PLN-NNNN)
      path: plan_id
      type: string
    - description: file-level slug with -XXXX suffix
      path: slug
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
slug: create-plan-t3mn
updated_at: "2026-04-14T03:27:04Z"
---
