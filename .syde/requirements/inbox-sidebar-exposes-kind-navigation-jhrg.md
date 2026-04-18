---
id: REQ-0231
kind: requirement
name: Inbox Sidebar Exposes Kind Navigation
slug: inbox-sidebar-exposes-kind-navigation-jhrg
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:37Z"
statement: When the user views any kind inbox screen, the dashboard shall display a kinds sidebar with entries for Systems, Components, Contracts, Concepts, Flows, and Decisions.
req_type: interface
priority: must
verification: manual inspection of any kind inbox in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Consistent sidebar navigation lets users switch between kinds without leaving the inbox pattern.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:37Z"
---
