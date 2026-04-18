---
id: REQ-0209
kind: requirement
name: Project Overview Returns Entity Counts Map
slug: project-overview-returns-entity-counts-map-arw5
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: http-api-afos
      type: refines
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:47Z"
statement: When GET /api/<project>/overview succeeds, the syded daemon shall return entity_counts as a map from entity kind to integer count.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The overview page tiles display per-kind entity totals.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:47Z"
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:47Z"
---
