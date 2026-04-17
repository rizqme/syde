---
id: CON-0028
kind: contract
name: Update Plan Phase
slug: update-plan-phase-izh0
description: Edit a plan phase's name, status, objective, changes, details, or notes.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan phase <plan-slug> <phase-id> [flags]
input_parameters:
    - path: plan-slug
      type: string
      description: positional, required
    - path: phase-id
      type: string
      description: positional, required. e.g. phase_1
    - path: --name
      type: string
      description: update name
    - path: --parent
      type: string
      description: change parent phase ID
    - path: --status
      type: string
      description: pending|in_progress|completed|skipped
    - path: --description
      type: string
      description: update description
    - path: --objective
      type: string
      description: update objective
    - path: --changes
      type: string
      description: update changes
    - path: --details
      type: string
      description: update details
    - path: --notes
      type: string
      description: update notes
output: exit 0; prints updated phase
output_parameters:
    - path: phase_id
      type: string
      description: echoed phase id
---
