---
id: REQ-0227
kind: requirement
name: Flows Inbox Route Renders List And Detail
slug: flows-inbox-route-renders-list-and-detail-1eij
relationships:
    - target: flows-inbox-screen-uh6s
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:36:50Z"
statement: When the user navigates to the /flow route, the dashboard shall render a two-column flows inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /flow in the dashboard
source: manual
source_ref: contract:flows-inbox-screen-uh6s
requirement_status: active
rationale: Flows inbox is the primary browsing surface for flow entities.
audited_overlaps:
    - slug: components-inbox-route-renders-list-and-detail-m6jv
      distinction: Targets the /flow route rendering flows; target covers the /component route rendering components.
    - slug: concepts-inbox-route-renders-list-mode-insk
      distinction: Flows route has no List toggle precondition; concepts route requires the List toggle active to render list mode.
    - slug: contracts-inbox-route-renders-list-and-detail-sjvw
      distinction: Targets the /flow route rendering flows; target covers the /contract route rendering contracts.
    - slug: decisions-inbox-route-renders-list-and-detail-eg0l
      distinction: Targets the /flow route rendering flows; target covers the /decision route rendering decisions.
    - slug: systems-inbox-route-renders-list-and-detail-bzae
      distinction: Targets the /flow route rendering flows; target covers the /system route rendering systems.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:50Z"
---
