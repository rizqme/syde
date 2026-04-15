---
id: CON-0068
kind: contract
name: Show Constraints
slug: show-constraints-jnig
description: Show all active decisions as project constraints.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: design-model-operations-coverage-wsrh
      type: involves
      label: flow
updated_at: "2026-04-15T09:17:19Z"
contract_kind: cli
interaction_pattern: request-response
input: syde constraints [--json]
input_parameters:
    - path: --json
      type: bool
      description: JSON output
output: All active decisions
output_parameters:
    - path: decisions
      type: '[]object'
      description: all architecture decisions
---
