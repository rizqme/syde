---
id: TSK-0116
kind: task
name: Build, reinstall skill, smoke test end-to-end gate
slug: build-reinstall-skill-smoke-test-end-to-end-gate-a7e7
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-review-strict-severity-verify-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: 'All layers work together: build clean, skill installed, gate fires, audit passes'
details: 'go build ./cmd/... && syde install-skill --all && syde reindex && syde sync check --strict. Then manual smoke: create a test requirement with a statement guaranteed to overlap an existing one and verify non-zero exit; rerun with --audited slug:reason and verify success. Clean up test req.'
acceptance: All commands exit 0; manual smoke confirms gate fires and succeeds via --audited
affected_entities:
    - audit-engine-4ktg
    - cli-commands-hpjb
    - skill-installer-wbmu
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_7
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:58:20Z"
---
