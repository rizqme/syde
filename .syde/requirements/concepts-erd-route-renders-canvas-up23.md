---
id: REQ-0224
kind: requirement
name: Concepts ERD Route Renders Canvas
slug: concepts-erd-route-renders-canvas-up23
relationships:
    - target: concepts-erd-screen-spxe
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:36Z"
statement: When the user navigates to the /concept route with the ERD toggle active, the dashboard shall render an ERD canvas containing every concept entity with attribute rows and relationship edges.
req_type: interface
priority: must
verification: manual inspection of /concept ERD mode in the dashboard
source: manual
source_ref: contract:concepts-erd-screen-spxe
requirement_status: active
rationale: ERD mode is the visual counterpart to the concept list view.
---
