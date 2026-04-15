---
contract_kind: storage
description: BadgerDB presence-only key for the (tag, kind, id) triple.
id: CON-0005
input: t:<tag>:<kind>:<id>
input_parameters:
    - description: tag label (lowercased)
      path: tag
      type: string
    - description: entity kind
      path: kind
      type: string
    - description: entity ID
      path: id
      type: string
interaction_pattern: schema
kind: contract
name: Tag Index Key
output: empty byte (presence-only)
output_parameters:
    - description: presence indicator; value is a single zero byte
      path: exists
      type: bool
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
slug: tag-index-key-5a4q
updated_at: "2026-04-14T11:06:21Z"
---
