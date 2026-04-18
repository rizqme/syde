---
id: CON-0086
kind: contract
name: Open Plan In Dashboard
slug: open-plan-in-dashboard-60th
description: CLI invocation that opens a plan in the syded dashboard, either by navigating an existing connected tab via WebSocket or by opening a new browser tab.
relationships:
- target: cli-commands
  type: references
- target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-17T10:45:59Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan open <plan-slug>
input_parameters:
- path: plan-slug
  type: string
  description: positional, required. Slug of the plan to open
output: exit 0 on success; stdout prints the dashboard URL
output_parameters:
- path: url
  type: string
  description: dashboard URL for the plan detail page
- path: mode
  type: string
  description: '''opened in existing dashboard tab'' or ''opened in new browser tab'''
---
