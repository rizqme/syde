---
id: REQ-0227
kind: requirement
name: Flows Inbox Route Renders List And Detail
slug: flows-inbox-route-renders-list-and-detail-1eij
relationships:
    - target: flows-inbox-screen-uh6s
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:41Z"
statement: When the user navigates to the /flow route, the dashboard shall render a two-column flows inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /flow in the dashboard
source: manual
source_ref: contract:flows-inbox-screen-uh6s
requirement_status: active
rationale: Flows inbox is the primary browsing surface for flow entities.
---
