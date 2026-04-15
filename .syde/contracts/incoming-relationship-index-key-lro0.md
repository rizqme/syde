---
contract_kind: storage
description: BadgerDB key listing all inbound relationships pointing at a target.
id: CON-0004
input: r:in:<target>
input_parameters:
    - description: target slug or ID
      path: target
      type: string
interaction_pattern: schema
kind: contract
name: Incoming Relationship Index Key
output: list of (source, type) pairs (gob-encoded)
output_parameters:
    - description: source entity ID
      path: source
      type: string
    - description: relationship type
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
slug: incoming-relationship-index-key-lro0
updated_at: "2026-04-14T11:06:21Z"
---
