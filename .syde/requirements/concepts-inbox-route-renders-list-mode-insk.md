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
updated_at: "2026-04-15T10:59:34Z"
statement: When the user navigates to the /concept route with the List toggle active, the dashboard shall render a two-column concepts inbox with a list panel and a detail panel.
req_type: interface
priority: must
verification: manual inspection of /concept in the dashboard
source: manual
source_ref: contract:concepts-inbox-screen-bpow
requirement_status: active
rationale: List mode is the default concept browsing experience.
---
