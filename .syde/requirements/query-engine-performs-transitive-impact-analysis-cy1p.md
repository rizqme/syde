---
id: REQ-0153
kind: requirement
name: Query Engine Performs Transitive Impact Analysis
slug: query-engine-performs-transitive-impact-analysis-cy1p
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:00Z"
statement: When Impacts is called for a slug, the query engine shall return the transitive set of entities that depend on the source.
req_type: functional
priority: must
verification: unit test of Impacts in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Impact analysis is critical for understanding change blast radius.
---
