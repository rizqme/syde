---
id: REQ-0216
kind: requirement
name: List Entities Accepts Kind Filter
slug: list-entities-accepts-kind-filter-rjf6
tags:
    - overlap-reviewed
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: http-api-afos
      type: refines
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:37:41Z"
statement: When GET /api/<project>/entities is invoked, the syded daemon shall accept an optional kind query parameter as a string and restrict results to that entity kind.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities?kind=component
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: Dashboard pages filter entity lists by kind.
audited_overlaps:
    - slug: list-entities-accepts-tag-filter-gdem
      distinction: Filters by entity kind query parameter; target filters by tag query parameter on the same endpoint.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:41Z"
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:41Z"
---
