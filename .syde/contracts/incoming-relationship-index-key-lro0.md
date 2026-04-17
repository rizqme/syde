---
id: CON-0004
kind: contract
name: Incoming Relationship Index Key
slug: incoming-relationship-index-key-lro0
description: BadgerDB key listing all inbound relationships pointing at a target.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: storage
interaction_pattern: schema
input: r:in:<target>
input_parameters:
    - path: target
      type: string
      description: target slug or ID
output: list of (source, type) pairs (gob-encoded)
output_parameters:
    - path: source
      type: string
      description: source entity ID
    - path: type
      type: string
      description: relationship type
---
