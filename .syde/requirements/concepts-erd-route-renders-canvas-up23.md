---
id: REQ-0224
kind: requirement
name: Concepts ERD Route Renders Canvas
slug: concepts-erd-route-renders-canvas-up23
relationships:
    - target: syded-dashboard-e82c
      type: belongs_to
    - target: erd-canvas-shall-be-removed
      type: refines
updated_at: "2026-04-17T11:03:29Z"
statement: When the user navigates to the /concept route with the ERD toggle active, the dashboard shall render an ERD canvas containing every concept entity with attribute rows and relationship edges.
req_type: interface
priority: must
verification: manual inspection of /concept ERD mode in the dashboard
source: manual
source_ref: contract:concepts-erd-screen-spxe
requirement_status: superseded
rationale: ERD mode is the visual counterpart to the concept list view.
superseded_by:
    - erd-canvas-shall-be-removed-4j8u
obsolete_reason: ERD canvas was removed from the concepts page in favor of glossary layout
---
