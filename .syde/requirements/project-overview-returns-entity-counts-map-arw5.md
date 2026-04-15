---
id: REQ-0209
kind: requirement
name: Project Overview Returns Entity Counts Map
slug: project-overview-returns-entity-counts-map-arw5
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:18Z"
statement: When GET /api/<project>/overview succeeds, the syded daemon shall return entity_counts as a map from entity kind to integer count.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The overview page tiles display per-kind entity totals.
---
