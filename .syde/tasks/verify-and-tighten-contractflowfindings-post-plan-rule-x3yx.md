---
id: TSK-0123
kind: task
name: Verify and tighten contractFlowFindings post-plan rule
slug: verify-and-tighten-contractflowfindings-post-plan-rule-x3yx
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every active contract is referenced by at least one active flow's steps; missing coverage is an ERROR with a clear message
details: Read internal/audit/graph_rules.go contractFlowFindings. Confirm it walks flow.Steps[].Contract (not legacy narrative scan) and emits ERROR (not WARN). Confirm the finding message names the missing contract. If any of these are off, fix. Add/refresh unit tests for the rule.
acceptance: 'go test ./internal/audit/... passes including contractFlowFindings; manual smoke: delete a step contract ref and confirm sync check errors with the contract''s name in the message'
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_11
created_at: "2026-04-17T09:50:47Z"
completed_at: "2026-04-17T10:23:56Z"
---
