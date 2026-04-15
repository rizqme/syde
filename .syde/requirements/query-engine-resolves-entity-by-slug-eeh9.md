---
id: REQ-0152
kind: requirement
name: Query Engine Resolves Entity By Slug
slug: query-engine-resolves-entity-by-slug-eeh9
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:58Z"
statement: When Lookup is called with a slug, the query engine shall resolve the entity by slug, bare-name, or ID and return the populated ResolvedEntity.
req_type: functional
priority: must
verification: unit test of Lookup in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Unified slug resolution is the read entry point for every query command.
---
