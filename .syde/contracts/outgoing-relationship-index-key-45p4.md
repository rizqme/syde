---
contract_kind: storage
description: BadgerDB key listing all outbound relationships from a source entity.
id: CON-0003
input: r:out:<id>
input_parameters:
    - description: source entity ID
      path: id
      type: string
interaction_pattern: schema
kind: contract
name: Outgoing Relationship Index Key
output: list of Relationship entries (gob-encoded)
output_parameters:
    - description: target slug or ID
      path: target
      type: string
    - description: relationship type e.g. belongs_to, depends_on, references
      path: type
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: outgoing-relationship-index-key-45p4
updated_at: "2026-04-14T11:06:21Z"
---
