---
id: FLW-0025
kind: flow
name: Browse Flows
slug: browse-flows-bnsc
description: User navigates to the flows inbox in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Flows in the sidebar
goal: User sees flow list and selects one to view its flowchart
steps:
    - id: s1
      action: User clicks Flows in sidebar
      contract: flows-inbox-screen
      description: Dashboard renders 2-column inbox with flowchart in detail
      on_success: done
---
