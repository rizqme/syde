---
id: REQ-0089
kind: requirement
name: SPA Is Read Only
slug: spa-is-read-only-0etl
relationships:
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:38Z"
statement: The web SPA shall not allow editing of any entity through the browser UI.
req_type: constraint
priority: must
verification: inspection of SPA routes confirming no mutation forms
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Writes go through the CLI; the dashboard stays a reviewer.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:38Z"
---
