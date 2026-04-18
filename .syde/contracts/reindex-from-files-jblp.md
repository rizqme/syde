---
id: CON-0021
kind: contract
name: Reindex From Files
slug: reindex-from-files-jblp
description: Rebuild the BadgerDB index from markdown files.
relationships:
- target: storage-engine
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde reindex
input_parameters:
- path: _
  type: '-'
  description: no arguments
output: Rebuilds BadgerDB index from all markdown files; prints stats
output_parameters:
- path: entities
  type: int
  description: number of entities indexed
- path: relationships
  type: int
  description: number of relationships indexed
- path: tags
  type: int
  description: number of tag entries
- path: words
  type: int
  description: number of word entries
---
