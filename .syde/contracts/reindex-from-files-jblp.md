---
contract_kind: cli
description: Rebuild the BadgerDB index from markdown files.
id: CON-0021
input: syde reindex
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Reindex From Files
output: Rebuilds BadgerDB index from all markdown files; prints stats
output_parameters:
    - description: number of entities indexed
      path: entities
      type: int
    - description: number of relationships indexed
      path: relationships
      type: int
    - description: number of tag entries
      path: tags
      type: int
    - description: number of word entries
      path: words
      type: int
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: reindex-from-files-jblp
updated_at: "2026-04-14T03:27:04Z"
---
