---
id: REQ-0154
kind: requirement
name: Query Engine Performs Relationship Traversal
slug: query-engine-performs-relationship-traversal-drni
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:03Z"
statement: When RelatedTo or DependsOn or DependedBy is called, the query engine shall return the entities connected to the source via the corresponding relationship directions.
req_type: functional
priority: must
verification: unit test of traversal helpers in engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Directional traversal feeds relationship-aware read commands.
---
