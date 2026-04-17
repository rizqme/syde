---
id: FLW-0029
kind: flow
name: Receive Live Updates
slug: receive-live-updates-pmid
description: Dashboard receives real-time entity changes via WebSocket
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: Entity is created, updated, or deleted in any client
goal: All connected dashboard tabs reflect the latest state
steps:
    - id: s1
      action: System broadcasts entity change
      contract: live-updates-websocket
      description: WebSocket pushes update event
      on_success: s2
    - id: s2
      action: Dashboard re-fetches affected data
      description: Client refreshes the view
      on_success: done
---
