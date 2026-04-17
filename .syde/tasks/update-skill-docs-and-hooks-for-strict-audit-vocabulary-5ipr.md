---
id: TSK-0128
kind: task
name: Update skill docs and hooks for strict audit vocabulary
slug: update-skill-docs-and-hooks-for-strict-audit-vocabulary-5ipr
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: SKILL.md, references, and hooks.json use 'Finding' vocabulary; every mention of 'warning' or '--strict' is gone or explicitly deprecated
details: 'grep skill/ and .claude/ for ''warning'', ''WARN '', ''--strict''. Rewrite each paragraph to reflect the new model: one severity level, all findings block completion. Reinstall via syde install-skill --all.'
acceptance: skill docs no longer reference --strict or warning; syde install-skill --all writes fresh copies
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/hooks.json
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_12
created_at: "2026-04-17T10:01:13Z"
completed_at: "2026-04-17T10:30:16Z"
---
