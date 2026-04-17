---
id: FLW-0021
kind: flow
name: Browse Components
slug: browse-components-b3f8
description: User navigates to the components inbox in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Components in the sidebar
goal: User sees component list and selects one for detail
steps:
    - id: s1
      action: User clicks Components in sidebar
      contract: components-inbox-screen
      description: Dashboard renders 2-column inbox
      on_success: done
---
