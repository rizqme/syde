---
id: CON-0020
kind: contract
name: Install Skill
slug: install-skill-bji4
description: Write skill files into .claude/ and append rules to CLAUDE.md.
relationships:
- target: skill-installer
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde install-skill
input_parameters:
- path: _
  type: '-'
  description: no arguments
output: Writes skill files into .claude/ and appends rules to CLAUDE.md
output_parameters:
- path: written
  type: '[]string'
  description: list of file paths written
---
