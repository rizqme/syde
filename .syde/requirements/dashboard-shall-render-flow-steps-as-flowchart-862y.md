---
id: REQ-0357
kind: requirement
name: Dashboard shall render flow steps as flowchart
slug: dashboard-shall-render-flow-steps-as-flowchart-862y
description: Visual flowchart with card nodes and directed edges
relationships:
    - target: web-spa
      type: refines
updated_at: "2026-04-18T09:37:19Z"
statement: When a flow entity has structured steps, the dashboard shall render them as a connected flowchart with card nodes and directed edges.
req_type: functional
priority: must
verification: Opening a flow with steps in dashboard shows a visual flowchart
source: plan
requirement_status: active
rationale: Prose flows are unreadable at scale; visual rendering is the payoff of structured steps
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:19Z"
---
