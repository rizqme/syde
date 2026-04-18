---
id: REQ-0155
kind: requirement
name: Query Engine Performs Full Text Search
slug: query-engine-performs-full-text-search-d9qd
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:37:42Z"
statement: When Search is called with a term, the query engine shall search across entity name, description, purpose, tags, and body fields.
req_type: functional
priority: must
verification: unit test of Search in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Text search is the primary fallback when slugs are unknown.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:37:42Z"
---
