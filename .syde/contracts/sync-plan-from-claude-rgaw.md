---
contract_kind: cli
description: Import a plan from a .claude/plans/*.md file.
id: CON-0034
input: syde plan sync [--file <path>]
input_parameters:
    - description: optional path to a .claude/plans/*.md file; defaults to most recent
      path: --file
      type: string
interaction_pattern: request-response
kind: contract
name: Sync Plan From Claude
output: Imports plan into syde and prints slug
output_parameters:
    - description: new plan slug
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
slug: sync-plan-from-claude-rgaw
updated_at: "2026-04-14T03:27:04Z"
---
