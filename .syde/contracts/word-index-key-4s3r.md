---
contract_kind: storage
description: BadgerDB key mapping a tokenized word to the entity field it appeared in.
id: CON-0006
input: w:<word>:<kind>:<id>
input_parameters:
    - description: tokenized, lowercased search term
      path: word
      type: string
    - description: entity kind
      path: kind
      type: string
    - description: entity ID
      path: id
      type: string
interaction_pattern: schema
kind: contract
name: Word Index Key
output: field name (string)
output_parameters:
    - description: which entity field this word came from (name|description|purpose|notes|body)
      path: field
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: word-index-key-4s3r
updated_at: "2026-04-14T11:06:21Z"
---
