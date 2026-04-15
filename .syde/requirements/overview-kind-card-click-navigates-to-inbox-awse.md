---
id: REQ-0207
kind: requirement
name: Overview Kind Card Click Navigates To Inbox
slug: overview-kind-card-click-navigates-to-inbox-awse
relationships:
    - target: overview-screen-2011
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:15Z"
statement: When the user clicks a kind card on the overview screen, the dashboard shall navigate to the corresponding kind inbox route.
req_type: interface
priority: should
verification: manual inspection of / in the dashboard
source: manual
source_ref: contract:overview-screen-2011
requirement_status: active
rationale: Kind cards are the primary drill-down affordance from the overview.
---
