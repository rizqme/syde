---
boundaries: Does NOT own entity data (delegates to storage). Does NOT perform impact analysis (query engine owns that).
capabilities:
    - Directed BFS over relationships with cycle detection
    - ASCII tree rendering for terminal display
    - Graphviz DOT export for visualization tools
description: Relationship-graph traversal and ASCII/DOT rendering.
files:
    - internal/graph/query.go
    - internal/graph/render.go
id: COM-0007
kind: component
name: Graph Engine
purpose: Compute and render the relationship graph between entities
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: depends_on
responsibility: BFS traversal + ASCII/DOT rendering for 'syde graph'
slug: graph-engine-xgjy
updated_at: "2026-04-14T03:35:54Z"
---
