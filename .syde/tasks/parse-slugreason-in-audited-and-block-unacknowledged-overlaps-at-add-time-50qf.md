---
id: TSK-0108
kind: task
name: Parse slug:reason in --audited and block unacknowledged overlaps at add time
slug: parse-slugreason-in-audited-and-block-unacknowledged-overlaps-at-add-time-50qf
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: syde add requirement with unacknowledged overlaps above 0.6 fails non-zero; --audited slug:reason acknowledges with distinction; --force bypasses
details: 'Edit internal/cli/add.go: after overlap detection, compare surfaced slugs against --audited entries. If any overlap slug is not covered and --force is not set, print the banner and return an error. Parse --audited values via strings.SplitN(v, '':'', 2) so slug-only entries still work (distinction empty). Update internal/cli/update.go similarly.'
acceptance: syde add requirement with a high-overlap statement and no --audited exits non-zero; with --audited slug:reason succeeds and writes distinction
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/add.go
    - internal/cli/update.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_2
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:13:52Z"
---
