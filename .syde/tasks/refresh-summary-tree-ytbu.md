---
id: TSK-0139
kind: task
name: Refresh summary tree
slug: refresh-summary-tree-ytbu
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T11:09:22Z"
task_status: completed
priority: medium
objective: syde tree status --strict exits 0.
details: syde tree scan, iterate syde tree changes --leaves-only --format json, dispatch subagents to summarize stale leaves, then summarize parent folders until clean.
acceptance: syde tree status --strict returns exit 0.
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_5
created_at: "2026-04-15T09:54:29Z"
completed_at: "2026-04-15T11:09:22Z"
---
