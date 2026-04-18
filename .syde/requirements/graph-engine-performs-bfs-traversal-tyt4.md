---
id: REQ-0146
kind: requirement
name: Graph Engine Performs BFS Traversal
slug: graph-engine-performs-bfs-traversal-tyt4
relationships:
    - target: graph-engine-xgjy
      type: refines
updated_at: "2026-04-18T09:37:20Z"
statement: When ImpactAnalysis is called with a source slug and max depth, the graph engine shall perform a directed BFS over relationships grouped by hop up to the given depth.
req_type: functional
priority: must
verification: unit test of ImpactAnalysis in internal/graph/query.go
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: BFS traversal is the core data structure for impact and neighborhood queries.
verified_against:
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:37:20Z"
---
