---
id: REQ-0357
kind: requirement
name: Dashboard shall render flow steps as flowchart
slug: dashboard-shall-render-flow-steps-as-flowchart-862y
description: Visual flowchart with card nodes and directed edges
relationships:
    - target: syde
      type: belongs_to
    - target: web-spa
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: When a flow entity has structured steps, the dashboard shall render them as a connected flowchart with card nodes and directed edges.
req_type: functional
priority: must
verification: Opening a flow with steps in dashboard shows a visual flowchart
source: plan
requirement_status: active
rationale: Prose flows are unreadable at scale; visual rendering is the payoff of structured steps
---
