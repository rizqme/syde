---
id: TSK-0112
kind: task
name: Review high-overlap template pairs (>=80 percent) and merge
slug: review-high-overlap-template-pairs-80-percent-and-merge-3xc4
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every >=80 percent similarity pair is resolved — merged into a survivor generic requirement, renamed to distinct wording, or re-acknowledged with a substantive distinction
details: 'Run go run ./cmd/listoverlaps. Walk the >=80 percent section. For each pair, syde query both, classify MERGE/RENAME/DISTINCT, execute. On MERGE: author a survivor requirement that refines the broader target, iterate inbound rels with two-step remove-then-add, mark originals superseded_by the survivor.'
acceptance: cmd/listoverlaps output shows zero unacknowledged pairs above 80 percent; acknowledged pairs carry non-empty distinction
affected_entities:
    - audit-engine-4ktg
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_5
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:57:26Z"
---
