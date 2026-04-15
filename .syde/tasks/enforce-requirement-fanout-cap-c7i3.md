---
id: TSK-0115
kind: task
name: Enforce requirement fanout cap
slug: enforce-requirement-fanout-cap-c7i3
relationships:
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: remove-learning-entity-and-cap-requirement-fanout-wqyz
      type: belongs_to
updated_at: "2026-04-15T09:23:33Z"
task_status: completed
priority: high
objective: Add an audit error when a requirement has more than 10 linked entities of the same kind.
details: Extend requirement traceability audit to count outbound links to requirement targets by requirement and source entity kind.
acceptance: sync strict reports baseline-style requirements that exceed the per-kind cap.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: remove-learning-entity-and-cap-requirement-fanout-wqyz
plan_phase: phase_3
created_at: "2026-04-15T08:23:05Z"
completed_at: "2026-04-15T09:09:24Z"
---
