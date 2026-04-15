---
contract_kind: cli
description: Scaffold .syde/ in the current project.
id: CON-0016
input: syde init [--install-skill]
input_parameters:
    - description: also run 'syde install-skill' after init
      path: --install-skill
      type: bool
interaction_pattern: request-response
kind: contract
name: Init Project
output: exit 0; creates .syde/ directory structure and syde.yaml
output_parameters:
    - description: list of created entity subdirectories
      path: created_dirs
      type: '[]string'
    - description: path to new syde.yaml
      path: config_path
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
slug: init-project-xm6c
updated_at: "2026-04-14T03:27:04Z"
---
