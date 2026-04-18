---
id: REQ-0219
kind: requirement
name: Contracts Inbox Route Renders List And Detail
slug: contracts-inbox-route-renders-list-and-detail-sjvw
relationships:
    - target: contracts-inbox-screen-x2tr
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:38:10Z"
statement: When the user navigates to the /contract route, the dashboard shall render a two-column contracts inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /contract in the dashboard
source: manual
source_ref: contract:contracts-inbox-screen-x2tr
requirement_status: active
rationale: Contracts inbox is the primary browsing surface for contract entities.
audited_overlaps:
    - slug: components-inbox-route-renders-list-and-detail-m6jv
      distinction: Targets the /contract route for contract entities; the components inbox requirement covers the /component route and component entities, a different URL and kind.
    - slug: concepts-inbox-route-renders-list-mode-insk
      distinction: Targets the /contract route with no mode toggle; the concepts inbox requirement governs the /concept route only when the List toggle is active.
    - slug: decisions-inbox-route-renders-list-and-detail-eg0l
      distinction: Targets the /contract route for contract entities; the decisions inbox requirement governs the /decision route for decision entities, a different URL and kind.
    - slug: flows-inbox-route-renders-list-and-detail-1eij
      distinction: Targets the /contract route and contract entities; the flows inbox requirement governs the /flow route and flow entities, a separate URL and entity kind.
    - slug: systems-inbox-route-renders-list-and-detail-bzae
      distinction: Targets the /contract route for contract entities; the systems inbox requirement governs the /system route for system entities, a different URL and kind.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:38:10Z"
---
