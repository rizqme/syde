---
id: CON-0007
kind: contract
name: Counter Key
slug: counter-key-4b8h
description: BadgerDB key holding the highest issued counter per entity kind.
relationships:
- target: storage-engine
  type: references
- type: belongs_to
  target: syde-5tdt
- type: belongs_to
  target: syded-dashboard-e82c
updated_at: '2026-04-16T10:51:15Z'
contract_kind: storage
interaction_pattern: schema
input: c:<kind>
input_parameters:
- path: kind
  type: string
  description: entity kind
output: ASCII decimal string of the highest issued counter
output_parameters:
- path: value
  type: int
  description: highest counter ever issued for this kind; NextID increments this atomically
---
