---
id: REQ-0155
kind: requirement
name: Query Engine Performs Full Text Search
slug: query-engine-performs-full-text-search-d9qd
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:05Z"
statement: When Search is called with a term, the query engine shall search across entity name, description, purpose, tags, and body fields.
req_type: functional
priority: must
verification: unit test of Search in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Text search is the primary fallback when slugs are unknown.
---
