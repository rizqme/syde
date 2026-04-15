---
contract_kind: cli
description: Modify an entity's fields, files, tags, or relationships.
id: CON-0011
input: syde update <slug> [flags]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: replace description
      path: --description
      type: string
    - description: replace purpose
      path: --purpose
      type: string
    - description: replace body
      path: --body
      type: string
    - description: repeatable 'target-slug:type'
      path: --add-rel
      type: '[]string'
    - description: repeatable target-slug to unlink
      path: --remove-rel
      type: '[]string'
    - description: add a tag
      path: --add-tag
      type: string
    - description: remove a tag
      path: --remove-tag
      type: string
    - description: replace files list
      path: --file
      type: '[]string'
    - description: append notes
      path: --note
      type: '[]string'
interaction_pattern: request-response
kind: contract
name: Update Entity
output: exit 0; prints updated slug
output_parameters:
    - description: echoed slug of updated entity
      path: slug
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
slug: update-entity-zpnh
updated_at: "2026-04-14T03:27:03Z"
---
