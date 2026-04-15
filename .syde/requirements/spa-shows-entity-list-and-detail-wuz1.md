---
id: REQ-0085
kind: requirement
name: SPA Shows Entity List And Detail
slug: spa-shows-entity-list-and-detail-wuz1
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The web SPA shall present kind-scoped entity list and detail views for every entity kind in the project.
req_type: functional
priority: must
verification: manual inspection of inbox routes for each entity kind
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Browsing entities by kind is the main reviewing workflow.
---
