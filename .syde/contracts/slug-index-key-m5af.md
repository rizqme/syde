---
id: CON-0002
kind: contract
name: Slug Index Key
slug: slug-index-key-m5af
description: BadgerDB key resolving a file slug back to its entity ID.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: storage
interaction_pattern: schema
input: s:<kind>:<slug>
input_parameters:
    - path: kind
      type: string
      description: entity kind
    - path: slug
      type: string
      description: file-level slug including the -XXXX suffix
output: entity ID (string)
output_parameters:
    - path: id
      type: string
      description: entity ID this slug resolves to
---
