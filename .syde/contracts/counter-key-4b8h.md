---
contract_kind: storage
description: BadgerDB key holding the highest issued counter per entity kind.
id: CON-0007
input: c:<kind>
input_parameters:
    - description: entity kind
      path: kind
      type: string
interaction_pattern: schema
kind: contract
name: Counter Key
output: ASCII decimal string of the highest issued counter
output_parameters:
    - description: highest counter ever issued for this kind; NextID increments this atomically
      path: value
      type: int
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: counter-key-4b8h
updated_at: "2026-04-14T11:06:21Z"
---
