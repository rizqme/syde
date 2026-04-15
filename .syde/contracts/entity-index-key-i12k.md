---
contract_kind: storage
description: BadgerDB key for the per-entity FileRef index.
id: CON-0001
input: e:<kind>:<id>
input_parameters:
    - description: entity kind (system|component|contract|...)
      path: kind
      type: string
    - description: counter-based entity ID e.g. COM-0001
      path: id
      type: string
interaction_pattern: schema
kind: contract
name: Entity Index Key
output: FileRef (gob-encoded)
output_parameters:
    - description: relative path to markdown file under .syde/<kind>/
      path: file
      type: string
    - description: entity ID
      path: id
      type: string
    - description: display name
      path: name
      type: string
    - description: entity kind
      path: kind
      type: string
    - description: field-name to [start,end] line range in the markdown file
      path: lines
      type: map<string,[2]int>
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
slug: entity-index-key-i12k
updated_at: "2026-04-14T11:06:21Z"
---
