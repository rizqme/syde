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
updated_at: "2026-04-16T10:51:16Z"
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
