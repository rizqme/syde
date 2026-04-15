---
id: REQ-0146
kind: requirement
name: Graph Engine Performs BFS Traversal
slug: graph-engine-performs-bfs-traversal-tyt4
relationships:
    - target: graph-engine-xgjy
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:42Z"
statement: When ImpactAnalysis is called with a source slug and max depth, the graph engine shall perform a directed BFS over relationships grouped by hop up to the given depth.
req_type: functional
priority: must
verification: unit test of ImpactAnalysis in internal/graph/query.go
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: BFS traversal is the core data structure for impact and neighborhood queries.
---
