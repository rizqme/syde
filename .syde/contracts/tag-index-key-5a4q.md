---
id: CON-0005
kind: contract
name: Tag Index Key
slug: tag-index-key-5a4q
description: BadgerDB presence-only key for the (tag, kind, id) triple.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: storage
interaction_pattern: schema
input: t:<tag>:<kind>:<id>
input_parameters:
    - path: tag
      type: string
      description: tag label (lowercased)
    - path: kind
      type: string
      description: entity kind
    - path: id
      type: string
      description: entity ID
output: empty byte (presence-only)
output_parameters:
    - path: exists
      type: bool
      description: presence indicator; value is a single zero byte
---
