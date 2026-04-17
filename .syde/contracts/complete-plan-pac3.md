---
id: CON-0085
kind: contract
name: Complete Plan
slug: complete-plan-pac3
description: CLI invocation that runs plan completion validation and marks the plan completed if every declared change matches actual entity state.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: audit-engine
      type: references
    - target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
      type: references
updated_at: "2026-04-17T10:45:59Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan complete <plan-slug> [--force]
input_parameters:
    - path: plan-slug
      type: string
      description: positional, required. Slug of the plan to complete
    - path: --force
      type: bool
      description: override audit errors (rare)
output: exit 0 on success; non-zero when sync check reports findings unless --force
output_parameters:
    - path: completed
      type: bool
      description: whether the plan was marked completed
    - path: findings
      type: array<Finding>
      description: plan completion findings blocking success
---
