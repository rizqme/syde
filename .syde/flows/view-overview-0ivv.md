---
id: FLW-0027
kind: flow
name: View Overview
slug: view-overview-0ivv
description: User views the project overview dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Overview in the sidebar or opens the dashboard root
goal: User sees entity counts, recent activity, and project health
steps:
    - id: s1
      action: User opens dashboard root
      contract: overview-screen
      description: Dashboard renders overview page
      on_success: s2
    - id: s2
      action: System loads project summary
      contract: project-overview
      description: API returns entity counts and health
      on_success: done
---
