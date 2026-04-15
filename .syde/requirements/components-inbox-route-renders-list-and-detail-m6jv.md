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
updated_at: "2026-04-15T10:59:24Z"
statement: When the user navigates to the /component route, the dashboard shall render a two-column components inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /component in the dashboard
source: manual
source_ref: contract:components-inbox-screen-c5jh
requirement_status: active
rationale: Components inbox is the primary browsing surface for component entities.
---
