---
id: TSK-0003
kind: task
name: Plan YAML round-trip smoke test
slug: plan-yaml-round-trip-smoke-test-3gp7
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: A plan with Design + one change per lane written to disk and read back matches byte-for-byte.
details: Hand-author a fixture plan markdown with a small Design, Extended component change with FieldChanges, New contract draft, Deleted requirement, then run syde reindex and syde query <plan-slug> --full --format json to confirm the Changes tree unmarshals cleanly.
acceptance: JSON output includes design + changes.components.extended[0].field_changes keys.
affected_entities:
    - query-engine-9k84
affected_files:
    - internal/query/formatter.go
    - internal/model/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_1
created_at: "2026-04-15T11:40:36Z"
completed_at: "2026-04-15T11:49:45Z"
---
