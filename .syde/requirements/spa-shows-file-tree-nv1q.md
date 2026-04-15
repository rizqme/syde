---
id: REQ-0087
kind: requirement
name: SPA Shows File Tree
slug: spa-shows-file-tree-nv1q
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The web SPA shall present a file tree explorer sourced from the project's summary tree.
req_type: functional
priority: must
verification: manual inspection of /tree route
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Users need to navigate the summary tree without leaving the dashboard.
---
