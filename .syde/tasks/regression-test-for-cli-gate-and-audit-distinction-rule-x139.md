---
id: TSK-0115
kind: task
name: Regression test for CLI gate and audit distinction rule
slug: regression-test-for-cli-gate-and-audit-distinction-rule-x139
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Go test asserts syde add requirement gate fires without --audited, succeeds with --audited slug:reason, and audit errors on empty distinction
details: 'Add internal/audit/requirements_test.go (or extend an existing test file) with table-driven cases: (a) high-overlap statement without --audited — expect error; (b) with --audited slug:rationale — expect success; (c) existing req with empty distinction — expect sync check error.'
acceptance: go test ./... passes; test covers all three cases
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/overlap_test.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_6
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:30:54Z"
---
