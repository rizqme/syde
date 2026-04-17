---
id: FLW-0024
kind: flow
name: Browse Systems
slug: browse-systems-ujn3
description: User navigates to the systems inbox in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Systems in the sidebar
goal: User sees system list and selects one for detail
steps:
    - id: s1
      action: User clicks Systems in sidebar
      contract: systems-inbox-screen
      description: Dashboard renders 2-column inbox
      on_success: done
---
