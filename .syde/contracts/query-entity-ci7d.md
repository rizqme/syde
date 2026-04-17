---
id: CON-0014
kind: contract
name: Query Entity
slug: query-entity-ci7d
description: Rich entity lookup with impact, traversal, search, diff.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: query-engine
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde query <slug> [flags]
input_parameters:
    - path: slug
      type: string
      description: positional, required
    - path: --full
      type: bool
      description: include body and all related data
    - path: --kind
      type: string
      description: filter by entity kind
    - path: --tag
      type: string
      description: filter by tag
    - path: --format
      type: string
      description: 'json|compact|refs (default: human)'
    - path: --impacts
      type: string
      description: transitive impact analysis starting at slug
    - path: --flow
      type: string
      description: flow decomposition
    - path: --components
      type: bool
      description: include component details in flow decomp
    - path: --related-to
      type: string
      description: direct connections for slug
    - path: --depends-on
      type: string
      description: entities this depends on
    - path: --depended-by
      type: string
      description: entities that depend on this
    - path: --search
      type: string
      description: full-text search
    - path: --diff
      type: string
      description: git change history
    - path: --since
      type: string
      description: time window for --diff (e.g. 7d)
output: formatted query result on stdout
output_parameters:
    - path: result
      type: string
      description: format depends on --format flag
---
