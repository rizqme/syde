---
id: REQ-0147
kind: requirement
name: Graph Engine Returns Direct Neighbors
slug: graph-engine-returns-direct-neighbors-9ski
relationships:
    - target: graph-engine-xgjy
      type: refines
updated_at: "2026-04-18T09:37:36Z"
statement: When Neighbors is called with an entity slug, the graph engine shall return the set of direct inbound and outbound edges for that entity.
req_type: functional
priority: must
verification: unit test of Neighbors in internal/graph/query.go
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: Neighborhood queries power the default syde graph output.
verified_against:
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:37:36Z"
---
