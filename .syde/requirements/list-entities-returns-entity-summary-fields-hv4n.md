---
id: REQ-0220
kind: requirement
name: List Entities Returns Entity Summary Fields
slug: list-entities-returns-entity-summary-fields-hv4n
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:31Z"
statement: When GET /api/<project>/entities succeeds, the syded daemon shall return each entity summary with id, slug, kind, and name as strings.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: These fields drive list rendering and navigation in the SPA.
---
