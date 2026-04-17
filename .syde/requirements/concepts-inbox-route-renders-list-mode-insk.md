---
id: REQ-0222
kind: requirement
name: Concepts Inbox Route Renders List Mode
slug: concepts-inbox-route-renders-list-mode-insk
relationships:
    - target: concepts-inbox-screen-bpow
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-17T10:45:44Z"
statement: When the user navigates to the /concept route with the List toggle active, the dashboard shall render a two-column concepts inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /concept in the dashboard
source: manual
source_ref: contract:concepts-inbox-screen-bpow
requirement_status: active
rationale: List mode is the default concept browsing experience.
audited_overlaps:
    - slug: components-inbox-route-renders-list-and-detail-m6jv
      distinction: Governs the /concept route when the List toggle is active; the components inbox requirement targets the /component route and has no list/graph toggle mode.
    - slug: contracts-inbox-route-renders-list-and-detail-sjvw
      distinction: Covers the /concept route in List mode; the contracts inbox requirement addresses the /contract route, a different URL and entity kind without a List-mode toggle.
    - slug: decisions-inbox-route-renders-list-and-detail-eg0l
      distinction: Covers the /concept route in List mode only; the decisions inbox requirement governs the /decision route for decision entities with no list/graph toggle.
    - slug: flows-inbox-route-renders-list-and-detail-1eij
      distinction: Governs the /concept route with List toggle active; the flows inbox requirement addresses the /flow route rendering flow entities and has no List-mode precondition.
    - slug: systems-inbox-route-renders-list-and-detail-bzae
      distinction: Covers the /concept route in List mode; the systems inbox requirement targets the /system route for system entities, a different URL without a list-mode toggle.
---
