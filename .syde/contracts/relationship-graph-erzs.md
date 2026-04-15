---
contract_kind: cli
description: Render the relationship graph as ASCII or Graphviz DOT.
id: CON-0023
input: syde graph [entity] [--format dot]
input_parameters:
    - description: positional, optional. Root entity; omit for global graph
      path: entity
      type: string
    - description: ascii (default) or dot (Graphviz)
      path: --format
      type: string
interaction_pattern: request-response
kind: contract
name: Relationship Graph
output: ASCII tree or DOT text on stdout
output_parameters:
    - description: rendered graph
      path: graph
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: graph-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: relationship-graph-erzs
updated_at: "2026-04-14T03:27:04Z"
---
