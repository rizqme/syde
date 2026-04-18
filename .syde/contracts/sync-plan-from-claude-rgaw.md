---
id: CON-0034
kind: contract
name: Sync Plan From Claude
slug: sync-plan-from-claude-rgaw
description: Import a plan from a .claude/plans/*.md file.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan sync [--file <path>]
input_parameters:
- path: --file
  type: string
  description: optional path to a .claude/plans/*.md file; defaults to most recent
output: Imports plan into syde and prints slug
output_parameters:
- path: slug
  type: string
  description: new plan slug
---
