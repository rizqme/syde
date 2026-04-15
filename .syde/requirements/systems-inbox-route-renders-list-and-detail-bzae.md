---
id: REQ-0208
kind: requirement
name: Systems Inbox Route Renders List And Detail
slug: systems-inbox-route-renders-list-and-detail-bzae
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:17Z"
statement: When the user navigates to the /system route, the dashboard shall render a two-column systems inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /system in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Systems inbox is the primary browsing surface for system entities.
---
