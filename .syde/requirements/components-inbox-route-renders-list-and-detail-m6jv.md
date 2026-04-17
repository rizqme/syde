---
id: REQ-0213
kind: requirement
name: Components Inbox Route Renders List And Detail
slug: components-inbox-route-renders-list-and-detail-m6jv
relationships:
    - target: components-inbox-screen-c5jh
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-17T10:46:08Z"
statement: When the user navigates to the /component route, the dashboard shall render a two-column components inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /component in the dashboard
source: manual
source_ref: contract:components-inbox-screen-c5jh
requirement_status: active
rationale: Components inbox is the primary browsing surface for component entities.
audited_overlaps:
    - slug: concepts-inbox-route-renders-list-mode-insk
      distinction: 'Different route, entity kind, and trigger: this fires on /component unconditionally, while the other fires on /concept only when the List toggle is active.'
    - slug: contracts-inbox-route-renders-list-and-detail-sjvw
      distinction: 'Different route and entity kind: this renders the /component inbox for components, while the other renders the /contract inbox for contracts.'
    - slug: decisions-inbox-route-renders-list-and-detail-eg0l
      distinction: 'Different route and entity kind: this renders the /component inbox for components, while the other renders the /decision inbox for decisions.'
    - slug: flows-inbox-route-renders-list-and-detail-1eij
      distinction: Targets the /component route rendering component entities; the flows inbox requirement governs the /flow route rendering flow entities, a distinct URL and entity kind.
    - slug: systems-inbox-route-renders-list-and-detail-bzae
      distinction: Covers the /component route and component list/detail; the systems inbox requirement covers the /system route and system entities, which are separate URL and entity kind.
---
