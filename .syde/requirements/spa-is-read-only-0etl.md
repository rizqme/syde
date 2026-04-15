---
id: REQ-0089
kind: requirement
name: SPA Is Read Only
slug: spa-is-read-only-0etl
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The web SPA shall not allow editing of any entity through the browser UI.
req_type: constraint
priority: must
verification: inspection of SPA routes confirming no mutation forms
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Writes go through the CLI; the dashboard stays a reviewer.
---
