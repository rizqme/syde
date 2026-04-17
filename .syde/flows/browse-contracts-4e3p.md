---
id: FLW-0022
kind: flow
name: Browse Contracts
slug: browse-contracts-4e3p
description: User navigates to the contracts inbox in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Contracts in the sidebar
goal: User sees contract list and selects one for detail
steps:
    - id: s1
      action: User clicks Contracts in sidebar
      contract: contracts-inbox-screen
      description: Dashboard renders 2-column inbox
      on_success: done
---
