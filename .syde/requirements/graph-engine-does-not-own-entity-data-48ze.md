---
id: REQ-0151
kind: requirement
name: Graph Engine Does Not Own Entity Data
slug: graph-engine-does-not-own-entity-data-48ze
relationships:
    - target: graph-engine-xgjy
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:53Z"
statement: The graph engine shall not own entity data and shall read relationships through the storage engine.
req_type: constraint
priority: must
verification: code review of internal/graph for direct persistence
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: Decoupling traversal from storage keeps the graph engine stateless and testable.
---
