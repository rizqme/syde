---
id: TSK-0253
kind: task
name: Simplify Graph.tsx to render every system at one tier
slug: simplify-graphtsx-to-render-every-system-at-one-tier-fv4c
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: dashboard-graph-shall-render-every-system-at-the-same-visual-tier-pb2g
      type: implements
updated_at: "2026-04-18T09:44:53Z"
task_status: completed
priority: high
objective: subSystemIds logic removed; every system kind uses the same sizeKey and radius; legend shows a single System entry.
acceptance: Dashboard /graph route renders every system at equal size with one legend entry labelled 'System'.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/pages/Graph.tsx
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_1
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:28:27Z"
---
