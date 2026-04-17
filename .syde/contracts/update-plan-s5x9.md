---
id: CON-0026
kind: contract
name: Update Plan
slug: update-plan-s5x9
description: Modify a plan's background, objective, scope, or description.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan update <slug> [flags]
input_parameters:
    - path: slug
      type: string
      description: positional, required
    - path: --background
      type: string
      description: update background
    - path: --objective
      type: string
      description: update objective
    - path: --scope
      type: string
      description: update scope
    - path: --description
      type: string
      description: update description
    - path: --purpose
      type: string
      description: update purpose
output: exit 0; prints updated slug
output_parameters:
    - path: slug
      type: string
      description: echoed slug
---
