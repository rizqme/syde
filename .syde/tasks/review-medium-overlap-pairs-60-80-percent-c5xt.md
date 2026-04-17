---
id: TSK-0113
kind: task
name: Review medium-overlap pairs (60-80 percent)
slug: review-medium-overlap-pairs-60-80-percent-c5xt
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every 60-80 percent similarity pair is resolved or re-acknowledged with documented distinction
details: Same workflow as high-overlap review but for the 60-80 percent band. Expect more RENAME and DISTINCT outcomes; merges are rarer at this tier.
acceptance: Zero unacknowledged pairs above 60 percent; acknowledged pairs carry non-empty distinction >=20 chars
affected_entities:
    - audit-engine-4ktg
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_5
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:57:26Z"
---
