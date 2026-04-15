---
acceptance: Sync check exits 0 after the loop
affected_entities:
    - cli-commands
    - skill-installer
affected_files:
    - skill/SKILL.md
    - internal/cli/task.go
completed_at: "2026-04-14T07:31:08Z"
created_at: "2026-04-14T07:24:22Z"
details: 'In a sandbox project: syde init; add a component with 3 files; plan + approve + task create with 1 affected entity; task start; touch all 3 files; syde task done <slug> --affected-entity <component> --affected-file f1 --affected-file f2 --affected-file f3; syde sync check --strict should exit 0.'
id: TSK-0035
kind: task
name: End-to-end smoke test
objective: A realistic task loop using the new flags leaves sync check clean
plan_phase: phase_2
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: end-to-end-smoke-test-3xhx
task_status: completed
updated_at: "2026-04-14T07:31:08Z"
---
