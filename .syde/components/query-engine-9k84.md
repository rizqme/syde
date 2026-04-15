---
boundaries: Does NOT mutate the store. Does NOT handle CLI flag parsing (CLI component owns that).
capabilities:
    - Transitive impact analysis (what breaks if this changes)
    - Relationship traversal (related-to, depends-on, depended-by)
    - Full-text search across name, description, purpose, body
    - Git diff over an entity's files
    - Output in human, json, compact, or refs format
description: 'Rich entity lookup engine: impact analysis, traversal, search, diff.'
files:
    - internal/query/engine.go
    - internal/query/resolver.go
    - internal/query/formatter.go
    - internal/query/diff.go
id: COM-0006
kind: component
name: Query Engine
purpose: Answer rich lookup questions about entities and their relationships
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: depends_on
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Resolve, format, and diff entities for the 'syde query' command
slug: query-engine-9k84
updated_at: "2026-04-15T06:26:49Z"
---
