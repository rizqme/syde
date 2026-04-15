---
id: REQ-0042
kind: requirement
name: HTTP API Routes Entity Endpoints
slug: http-api-routes-entity-endpoints-8esf
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: The syded HTTP API shall handle HTTP routes for entities, tree, search, constraints, and context requests.
req_type: functional
priority: must
verification: integration test against /api/*/entities and related endpoints
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Routing HTTP requests to entity data is the component's core responsibility.
---
