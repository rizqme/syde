---
id: CON-0023
kind: contract
name: Relationship Graph
slug: relationship-graph-erzs
description: Render the relationship graph as ASCII or Graphviz DOT.
relationships:
- target: graph-engine
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde graph [entity] [--format dot]
input_parameters:
- path: entity
  type: string
  description: positional, optional. Root entity; omit for global graph
- path: --format
  type: string
  description: ascii (default) or dot (Graphviz)
output: ASCII tree or DOT text on stdout
output_parameters:
- path: graph
  type: string
  description: rendered graph
---
