---
id: REQ-0147
kind: requirement
name: Graph Engine Returns Direct Neighbors
slug: graph-engine-returns-direct-neighbors-9ski
relationships:
    - target: graph-engine-xgjy
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:44Z"
statement: When Neighbors is called with an entity slug, the graph engine shall return the set of direct inbound and outbound edges for that entity.
req_type: functional
priority: must
verification: unit test of Neighbors in internal/graph/query.go
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: Neighborhood queries power the default syde graph output.
---
