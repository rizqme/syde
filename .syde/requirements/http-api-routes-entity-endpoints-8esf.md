---
id: REQ-0042
kind: requirement
name: HTTP API Routes Entity Endpoints
slug: http-api-routes-entity-endpoints-8esf
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:44Z"
statement: The syded HTTP API shall handle HTTP routes for entities, tree, search, constraints, and context requests.
req_type: functional
priority: must
verification: integration test against /api/*/entities and related endpoints
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Routing HTTP requests to entity data is the component's core responsibility.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:44Z"
---
