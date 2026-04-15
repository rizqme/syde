---
contract_kind: cli
description: Rich entity lookup with impact, traversal, search, diff.
id: CON-0014
input: syde query <slug> [flags]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: include body and all related data
      path: --full
      type: bool
    - description: filter by entity kind
      path: --kind
      type: string
    - description: filter by tag
      path: --tag
      type: string
    - description: 'json|compact|refs (default: human)'
      path: --format
      type: string
    - description: transitive impact analysis starting at slug
      path: --impacts
      type: string
    - description: flow decomposition
      path: --flow
      type: string
    - description: include component details in flow decomp
      path: --components
      type: bool
    - description: direct connections for slug
      path: --related-to
      type: string
    - description: entities this depends on
      path: --depends-on
      type: string
    - description: entities that depend on this
      path: --depended-by
      type: string
    - description: full-text search
      path: --search
      type: string
    - description: git change history
      path: --diff
      type: string
    - description: time window for --diff (e.g. 7d)
      path: --since
      type: string
interaction_pattern: request-response
kind: contract
name: Query Entity
output: formatted query result on stdout
output_parameters:
    - description: format depends on --format flag
      path: result
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: query-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: query-entity-ci7d
updated_at: "2026-04-14T03:27:04Z"
---
