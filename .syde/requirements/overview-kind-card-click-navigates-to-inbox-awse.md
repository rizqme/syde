---
id: REQ-0207
kind: requirement
name: Overview Kind Card Click Navigates To Inbox
slug: overview-kind-card-click-navigates-to-inbox-awse
relationships:
    - target: overview-screen-2011
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:38:09Z"
statement: When the user clicks a kind card on the overview screen, the dashboard shall navigate to the corresponding kind inbox route.
req_type: interface
priority: should
verification: manual inspection of / in the dashboard
source: manual
source_ref: contract:overview-screen-2011
requirement_status: active
rationale: Kind cards are the primary drill-down affordance from the overview.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:38:09Z"
---
