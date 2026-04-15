---
id: REQ-0219
kind: requirement
name: Contracts Inbox Route Renders List And Detail
slug: contracts-inbox-route-renders-list-and-detail-sjvw
relationships:
    - target: contracts-inbox-screen-x2tr
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:29Z"
statement: When the user navigates to the /contract route, the dashboard shall render a two-column contracts inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /contract in the dashboard
source: manual
source_ref: contract:contracts-inbox-screen-x2tr
requirement_status: active
rationale: Contracts inbox is the primary browsing surface for contract entities.
---
