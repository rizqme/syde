---
id: REQ-0231
kind: requirement
name: Inbox Sidebar Exposes Kind Navigation
slug: inbox-sidebar-exposes-kind-navigation-jhrg
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:45Z"
statement: When the user views any kind inbox screen, the dashboard shall display a kinds sidebar with entries for Systems, Components, Contracts, Concepts, Flows, and Decisions.
req_type: interface
priority: must
verification: manual inspection of any kind inbox in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Consistent sidebar navigation lets users switch between kinds without leaving the inbox pattern.
---
