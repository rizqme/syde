---
id: REQ-0090
kind: requirement
name: SPA Talks Only Via HTTP API
slug: spa-talks-only-via-http-api-c91t
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The web SPA shall not access BadgerDB or .syde markdown files directly and shall fetch all data through the HTTP API.
req_type: constraint
priority: must
verification: inspection of web/src for any non-HTTP data access
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Routing data through HTTP preserves the single-writer property of the store.
---
