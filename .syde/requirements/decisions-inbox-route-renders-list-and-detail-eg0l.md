---
id: REQ-0229
kind: requirement
name: Decisions Inbox Route Renders List And Detail
slug: decisions-inbox-route-renders-list-and-detail-eg0l
relationships:
    - target: decisions-inbox-screen-wnsc
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:42Z"
statement: When the user navigates to the /decision route, the dashboard shall render a two-column decisions inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /decision in the dashboard
source: manual
source_ref: contract:decisions-inbox-screen-wnsc
requirement_status: active
rationale: Decisions inbox is the primary browsing surface for decision entities.
---
