---
id: REQ-0225
kind: requirement
name: Get Entity Accepts Full Or Bare Slug
slug: get-entity-accepts-full-or-bare-slug-qh8d
relationships:
    - target: get-entity-http-0949
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:39Z"
statement: When GET /api/<project>/entity/<slug> is invoked, the syded daemon shall accept the slug path parameter as a string in either full or bare form and resolve it to a single entity.
req_type: interface
priority: must
verification: integration test against /api/<project>/entity/<slug>
source: manual
source_ref: contract:get-entity-http-0949
requirement_status: active
rationale: Callers may pass either the full slug or the bare name portion.
---
