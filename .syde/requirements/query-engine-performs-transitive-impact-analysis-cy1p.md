---
id: REQ-0153
kind: requirement
name: Query Engine Performs Transitive Impact Analysis
slug: query-engine-performs-transitive-impact-analysis-cy1p
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:37:00Z"
statement: When Impacts is called for a slug, the query engine shall return the transitive set of entities that depend on the source.
req_type: functional
priority: must
verification: unit test of Impacts in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Impact analysis is critical for understanding change blast radius.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:37:00Z"
---
