---
id: FLW-0031
kind: flow
name: Open Dashboard
slug: open-dashboard-3nka
description: User opens the dashboard in a browser
tags:
    - init
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User runs syde open or clicks a dashboard link
goal: Dashboard is visible in the browser showing the project
steps:
    - id: s1
      action: User runs syde open
      contract: open-dashboard
      description: Spawns syded if needed, opens browser
      on_success: s2
    - id: s2
      action: System lists registered projects
      contract: list-projects
      description: Returns project slugs
      on_success: done
---
