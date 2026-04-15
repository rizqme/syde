---
contract_kind: cli
description: Render a plan as a hierarchical tree with optional --full detail.
id: CON-0031
input: syde plan show <slug> [--full]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: include phase details and per-task objective/details/acceptance
      path: --full
      type: bool
interaction_pattern: request-response
kind: contract
name: Show Plan
output: ASCII tree of plan phases and tasks
output_parameters:
    - description: rendered hierarchical plan
      path: tree
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: show-plan-1ybx
updated_at: "2026-04-14T03:27:04Z"
---
