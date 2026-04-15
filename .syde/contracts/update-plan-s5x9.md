---
contract_kind: cli
description: Modify a plan's background, objective, scope, or description.
id: CON-0026
input: syde plan update <slug> [flags]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: update background
      path: --background
      type: string
    - description: update objective
      path: --objective
      type: string
    - description: update scope
      path: --scope
      type: string
    - description: update description
      path: --description
      type: string
    - description: update purpose
      path: --purpose
      type: string
interaction_pattern: request-response
kind: contract
name: Update Plan
output: exit 0; prints updated slug
output_parameters:
    - description: echoed slug
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
slug: update-plan-s5x9
updated_at: "2026-04-14T03:27:04Z"
---
