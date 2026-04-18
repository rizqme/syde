---
id: REQ-0154
kind: requirement
name: Query Engine Performs Relationship Traversal
slug: query-engine-performs-relationship-traversal-drni
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:38:01Z"
statement: When RelatedTo or DependsOn or DependedBy is called, the query engine shall return the entities connected to the source via the corresponding relationship directions.
req_type: functional
priority: must
verification: unit test of traversal helpers in engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Directional traversal feeds relationship-aware read commands.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:38:01Z"
---
