---
id: REQ-0090
kind: requirement
name: SPA Talks Only Via HTTP API
slug: spa-talks-only-via-http-api-c91t
relationships:
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:09Z"
statement: The web SPA shall not access BadgerDB or .syde markdown files directly and shall fetch all data through the HTTP API.
req_type: constraint
priority: must
verification: inspection of web/src for any non-HTTP data access
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Routing data through HTTP preserves the single-writer property of the store.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:09Z"
---
