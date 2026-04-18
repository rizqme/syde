---
id: REQ-0208
kind: requirement
name: Systems Inbox Route Renders List And Detail
slug: systems-inbox-route-renders-list-and-detail-bzae
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:35Z"
statement: When the user navigates to the /system route, the dashboard shall render a two-column systems inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /system in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Systems inbox is the primary browsing surface for system entities.
audited_overlaps:
    - slug: components-inbox-route-renders-list-and-detail-m6jv
      distinction: The /system route renders the systems inbox; the /component route renders the components inbox, different URLs and entity kinds.
    - slug: concepts-inbox-route-renders-list-mode-insk
      distinction: /system always renders the two-column inbox; /concept only renders list mode when the List toggle is active, different routes and trigger conditions.
    - slug: contracts-inbox-route-renders-list-and-detail-sjvw
      distinction: /system route renders systems inbox; /contract route renders contracts inbox, different URLs and entity kinds.
    - slug: decisions-inbox-route-renders-list-and-detail-eg0l
      distinction: /system route renders systems inbox; /decision route renders decisions inbox, different URLs and entity kinds.
    - slug: flows-inbox-route-renders-list-and-detail-1eij
      distinction: /system route renders systems inbox; /flow route renders flows inbox, different URLs and entity kinds.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:35Z"
---
