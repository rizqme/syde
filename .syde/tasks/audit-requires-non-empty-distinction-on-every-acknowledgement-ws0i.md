---
id: TSK-0109
kind: task
name: Audit requires non-empty distinction on every acknowledgement
slug: audit-requires-non-empty-distinction-on-every-acknowledgement-ws0i
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: sync check errors on audited_overlaps entries where distinction is empty or shorter than 20 characters
details: 'Edit internal/audit/requirements.go: extend requirementOverlapFindings (or add a new rule) to walk audited_overlaps on every active requirement and emit ERROR when Distinction is empty or below a minimum length. Keep the existing unacknowledged-overlap finding.'
acceptance: syde sync check reports errors for every existing acknowledgement with empty distinction (used as test fixtures for Phase 5)
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/requirements.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_3
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:15:33Z"
---
