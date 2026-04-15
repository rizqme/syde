---
acceptance: Hook fires; orphan in repo blocks session end
affected_entities:
    - skill-installer
affected_files:
    - skill/hooks.json
    - skill/SKILL.md
completed_at: "2026-04-14T06:18:43Z"
created_at: "2026-04-14T06:03:22Z"
details: skill/hooks.json Stop hook runs syde sync check --strict; skill/SKILL.md Phase 5 rewritten to reference the single command.
id: TSK-0003
kind: task
name: Wire Stop hook to syde sync check --strict
objective: Session end is blocked when audit not clean
plan_phase: phase_2
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: wire-stop-hook-to-syde-sync-check-strict-31vs
task_status: completed
updated_at: "2026-04-14T06:18:43Z"
---
