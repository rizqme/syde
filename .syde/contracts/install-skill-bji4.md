---
contract_kind: cli
description: Write skill files into .claude/ and append rules to CLAUDE.md.
id: CON-0020
input: syde install-skill
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Install Skill
output: Writes skill files into .claude/ and appends rules to CLAUDE.md
output_parameters:
    - description: list of file paths written
      path: written
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: skill-installer
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: install-skill-bji4
updated_at: "2026-04-14T03:27:04Z"
---
