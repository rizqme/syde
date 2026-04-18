---
id: CON-0087
kind: contract
name: Sync Check Gate
slug: sync-check-gate-fjd6
description: Canonical session-end health gate. Runs the full audit and exits non-zero on any finding — there is no non-strict mode.
relationships:
- target: cli-commands
  type: references
- target: audit-engine
  type: references
- target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-17T10:45:59Z'
contract_kind: cli
interaction_pattern: request-response
input: syde sync check
input_parameters:
- path: --format
  type: string
  description: output format (rich, json). Default rich
output: exit 0 on clean; non-zero on any finding
output_parameters:
- path: findings
  type: array<Finding>
  description: all findings across audit, tree, orphan, drift, overlap, coverage rules
- path: entities
  type: int
  description: total entity count scanned
---
