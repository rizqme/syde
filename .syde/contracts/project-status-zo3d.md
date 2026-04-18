---
id: CON-0018
kind: contract
name: Project Status
slug: project-status-zo3d
description: Show entity counts by kind.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde status
input_parameters:
- path: _
  type: '-'
  description: no arguments
output: Entity counts by kind on stdout
output_parameters:
- path: counts
  type: map<string,int>
  description: kind → count
---
