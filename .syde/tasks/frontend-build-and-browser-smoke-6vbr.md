---
id: TSK-0032
kind: task
name: Frontend build and browser smoke
slug: frontend-build-and-browser-smoke-6vbr
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: pending
priority: high
objective: bun run build passes; opening /<proj>/plan in the browser shows the 2-column layout with the new plan in the list and the PlanDetailPanel content on the right.
details: cd web && bun run build. make install. Open the dashboard, navigate to /<proj>/plan, click the Plans Inbox 2-Column Layout plan, switch between Plan and Tasks tabs to confirm both render. Verify no floating panel appears.
acceptance: Build clean; manual browser test passes; no console errors.
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_5
created_at: "2026-04-15T13:03:56Z"
---
