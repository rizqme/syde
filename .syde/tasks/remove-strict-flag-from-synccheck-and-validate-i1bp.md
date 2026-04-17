---
id: TSK-0127
kind: task
name: Remove --strict flag from sync_check and validate
slug: remove-strict-flag-from-synccheck-and-validate-i1bp
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-review-strict-severity-verify-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: syde sync check and syde validate no longer accept --strict; stale-tree cascade always blocks
details: 'Edit internal/cli/sync_check.go: drop the --strict flag binding and the conditional paths that used it. If syde validate has --strict, do the same. Update command help text.'
acceptance: syde sync check --strict exits non-zero with unknown flag error; syde sync check without any flag blocks on stale tree
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/sync_check.go
    - internal/cli/validate.go
    - internal/cli/codex_hook.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_12
created_at: "2026-04-17T10:01:13Z"
completed_at: "2026-04-17T10:27:44Z"
---
