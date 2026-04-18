---
id: CON-0011
kind: contract
name: Update Entity
slug: update-entity-zpnh
description: Modify an entity's fields, files, tags, or relationships.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde update <slug> [flags]
input_parameters:
- path: slug
  type: string
  description: positional, required
- path: --description
  type: string
  description: replace description
- path: --purpose
  type: string
  description: replace purpose
- path: --body
  type: string
  description: replace body
- path: --add-rel
  type: '[]string'
  description: repeatable 'target-slug:type'
- path: --remove-rel
  type: '[]string'
  description: repeatable target-slug to unlink
- path: --add-tag
  type: string
  description: add a tag
- path: --remove-tag
  type: string
  description: remove a tag
- path: --file
  type: '[]string'
  description: replace files list
- path: --note
  type: '[]string'
  description: append notes
output: exit 0; prints updated slug
output_parameters:
- path: slug
  type: string
  description: echoed slug of updated entity
---
