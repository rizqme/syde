---
id: TSK-0017
kind: task
name: Frontend build and browser smoke
slug: frontend-build-and-browser-smoke-xiqx
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: bun run build succeeds and a dev instance shows a plan with mixed changes correctly.
details: cd web && bun run build. Then fire up syded and manually open a test plan in Chrome to confirm layout, kind switching, and specialized contract views render without console errors. Update any broken snapshot in the build.
acceptance: bun run build green, manual browser inspection clean.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:19:30Z"
---
