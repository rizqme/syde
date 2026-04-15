---
contract_kind: storage
description: BadgerDB key resolving a file slug back to its entity ID.
id: CON-0002
input: s:<kind>:<slug>
input_parameters:
    - description: entity kind
      path: kind
      type: string
    - description: file-level slug including the -XXXX suffix
      path: slug
      type: string
interaction_pattern: schema
kind: contract
name: Slug Index Key
output: entity ID (string)
output_parameters:
    - description: entity ID this slug resolves to
      path: id
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: slug-index-key-m5af
updated_at: "2026-04-14T11:06:21Z"
---
