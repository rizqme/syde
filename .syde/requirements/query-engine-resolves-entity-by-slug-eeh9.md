---
id: REQ-0152
kind: requirement
name: Query Engine Resolves Entity By Slug
slug: query-engine-resolves-entity-by-slug-eeh9
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:37:03Z"
statement: When Lookup is called with a slug, the query engine shall resolve the entity by slug, bare-name, or ID and return the populated ResolvedEntity.
req_type: functional
priority: must
verification: unit test of Lookup in internal/query/engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Unified slug resolution is the read entry point for every query command.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:37:03Z"
---
