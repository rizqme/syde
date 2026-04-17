---
id: FLW-0026
kind: flow
name: Browse Decisions
slug: browse-decisions-vh6o
description: User navigates to the decisions inbox in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Decisions in the sidebar
goal: User sees decision list and selects one for detail
steps:
    - id: s1
      action: User clicks Decisions in sidebar
      contract: decisions-inbox-screen
      description: Dashboard renders 2-column inbox
      on_success: done
---
