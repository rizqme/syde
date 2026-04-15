---
id: COM-0006
kind: component
name: Query Engine
slug: query-engine-9k84
description: 'Rich entity lookup engine: impact analysis, traversal, search, diff.'
purpose: Answer rich lookup questions about entities and their relationships
files:
    - internal/query/engine.go
    - internal/query/resolver.go
    - internal/query/formatter.go
    - internal/query/diff.go
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: depends_on
updated_at: "2026-04-15T10:58:06Z"
responsibility: Resolve, format, and diff entities for the 'syde query' command
capabilities:
    - Transitive impact analysis (what breaks if this changes)
    - Relationship traversal (related-to, depends-on, depended-by)
    - Full-text search across name, description, purpose, body
    - Git diff over an entity's files
    - Output in human, json, compact, or refs format
boundaries: Does NOT mutate the store. Does NOT handle CLI flag parsing (CLI component owns that).
---
