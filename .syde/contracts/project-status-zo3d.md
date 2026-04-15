---
contract_kind: cli
description: Show entity counts by kind.
id: CON-0018
input: syde status
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Project Status
output: Entity counts by kind on stdout
output_parameters:
    - description: kind → count
      path: counts
      type: map<string,int>
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: project-status-zo3d
updated_at: "2026-04-14T03:27:04Z"
---
