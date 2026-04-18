---
id: CON-0006
kind: contract
name: Word Index Key
slug: word-index-key-4s3r
description: BadgerDB key mapping a tokenized word to the entity field it appeared in.
relationships:
- target: storage-engine
  type: references
- type: belongs_to
  target: syde-5tdt
- type: belongs_to
  target: syded-dashboard-e82c
updated_at: '2026-04-16T10:51:16Z'
contract_kind: storage
interaction_pattern: schema
input: w:<word>:<kind>:<id>
input_parameters:
- path: word
  type: string
  description: tokenized, lowercased search term
- path: kind
  type: string
  description: entity kind
- path: id
  type: string
  description: entity ID
output: field name (string)
output_parameters:
- path: field
  type: string
  description: which entity field this word came from (name|description|purpose|notes|body)
---
