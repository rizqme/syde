---
id: TSK-0121
kind: task
name: Build parity registry and symmetry test
slug: build-parity-registry-and-symmetry-test-40bb
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every plan_authoring rule has an equivalent post-plan rule encoded in a parity registry and asserted by a Go test
details: Create internal/audit/symmetry.go with SymmetryEntry{PlanningRule, PostPlanRule, Description} slice covering every current plan_authoring rule. Create symmetry_test.go that iterates the slice and, for each entry, (a) asserts a planning finding of the given category+field is produced by planAuthoringFindings on a crafted input, and (b) asserts the post-plan counterpart fires on the same input at sync check. Fill any missing post-plan counterpart in requirements.go or graph_rules.go. Document in skill/SKILL.md.
acceptance: go test ./internal/audit/... passes including the new symmetry test; SKILL.md documents the principle
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/symmetry.go
    - internal/audit/symmetry_test.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_10
created_at: "2026-04-17T09:46:36Z"
completed_at: "2026-04-17T10:28:46Z"
---
