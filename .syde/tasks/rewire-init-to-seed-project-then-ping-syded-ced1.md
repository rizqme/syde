---
acceptance: syde init in a fresh dir creates .syde/ and the project is visible on GET /api/projects without manual restart
completed_at: "2026-04-14T06:58:54Z"
details: init writes files directly (no Store needed), then calls client.New (auto-launching syded) so the new project appears in the dashboard immediately.
id: TSK-0026
kind: task
name: Rewire init to seed project then ping syded
objective: syde init creates .syde/ on disk and registers the project with syded
plan_phase: phase_4
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-init-to-seed-project-then-ping-syded-ced1
task_status: completed
updated_at: "2026-04-14T06:58:54Z"
---
