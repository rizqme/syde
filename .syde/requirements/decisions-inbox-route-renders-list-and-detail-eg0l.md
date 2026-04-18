---
id: REQ-0229
kind: requirement
name: Decisions Inbox Route Renders List And Detail
slug: decisions-inbox-route-renders-list-and-detail-eg0l
relationships:
    - target: decisions-inbox-screen-wnsc
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:25Z"
statement: When the user navigates to the /decision route, the dashboard shall render a two-column decisions inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /decision in the dashboard
source: manual
source_ref: contract:decisions-inbox-screen-wnsc
requirement_status: active
rationale: Decisions inbox is the primary browsing surface for decision entities.
audited_overlaps:
    - slug: components-inbox-route-renders-list-and-detail-m6jv
      distinction: Targets the /decision route for decision entities; the components inbox requirement governs the /component route for component entities, a different URL and entity kind.
    - slug: concepts-inbox-route-renders-list-mode-insk
      distinction: Targets the /decision route unconditionally; the concepts inbox requirement governs the /concept route only when the List toggle is active and covers concept entities instead.
    - slug: contracts-inbox-route-renders-list-and-detail-sjvw
      distinction: Covers the /decision route and the decisions entity kind, not the /contract route or contract entities.
    - slug: flows-inbox-route-renders-list-and-detail-1eij
      distinction: Covers the /decision route and the decisions entity kind, not the /flow route or flow entities.
    - slug: systems-inbox-route-renders-list-and-detail-bzae
      distinction: Covers the /decision route and the decisions entity kind, not the /system route or system entities.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:25Z"
---
