---
id: CON-0001
kind: contract
name: Entity Index Key
slug: entity-index-key-i12k
description: BadgerDB key for the per-entity FileRef index.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: storage
interaction_pattern: schema
input: e:<kind>:<id>
input_parameters:
    - path: kind
      type: string
      description: entity kind (system|component|contract|...)
    - path: id
      type: string
      description: counter-based entity ID e.g. COM-0001
output: FileRef (gob-encoded)
output_parameters:
    - path: file
      type: string
      description: relative path to markdown file under .syde/<kind>/
    - path: id
      type: string
      description: entity ID
    - path: name
      type: string
      description: display name
    - path: kind
      type: string
      description: entity kind
    - path: lines
      type: map<string,[2]int>
      description: field-name to [start,end] line range in the markdown file
---
