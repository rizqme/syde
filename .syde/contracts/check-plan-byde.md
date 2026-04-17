---
id: CON-0084
kind: contract
name: Check Plan
slug: check-plan-byde
description: CLI invocation that runs the plan authoring audit and prints every finding; used as a pre-approval gate.
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
input: syde plan check <plan-slug>
input_parameters:
    - path: plan-slug
      type: string
      description: positional, required. Slug of the plan to audit
output: exit 0 on clean check; non-zero when any finding is reported
output_parameters:
    - path: findings
      type: array<Finding>
      description: plan_authoring findings printed to stdout
---
