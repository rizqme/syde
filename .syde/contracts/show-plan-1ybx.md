---
id: CON-0031
kind: contract
name: Show Plan
slug: show-plan-1ybx
description: Render a plan as a hierarchical tree with optional --full detail.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan show <slug> [--full]
input_parameters:
    - path: slug
      type: string
      description: positional, required
    - path: --full
      type: bool
      description: include phase details and per-task objective/details/acceptance
output: ASCII tree of plan phases and tasks
output_parameters:
    - path: tree
      type: string
      description: rendered hierarchical plan
---
