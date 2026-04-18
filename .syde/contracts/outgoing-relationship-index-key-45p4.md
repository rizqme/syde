---
id: CON-0003
kind: contract
name: Outgoing Relationship Index Key
slug: outgoing-relationship-index-key-45p4
description: BadgerDB key listing all outbound relationships from a source entity.
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
input: r:out:<id>
input_parameters:
- path: id
  type: string
  description: source entity ID
output: list of Relationship entries (gob-encoded)
output_parameters:
- path: target
  type: string
  description: target slug or ID
- path: type
  type: string
  description: relationship type e.g. belongs_to, depends_on, references
---
