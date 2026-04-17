---
id: TSK-0126
kind: task
name: Reclassify every Warning and Hint producer as Finding
slug: reclassify-every-warning-and-hint-producer-as-finding-ab9d
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-review-strict-severity-verify-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every SeverityWarning and SeverityHint call site in internal/audit now emits Finding; no severity downgrades remain
details: grep every file under internal/audit/ for SeverityWarning and SeverityHint. For each, change to the new single Finding severity. Spot-check plan_authoring, requirements, graph_rules, orphans, screens, plan_completion.
acceptance: grep returns zero SeverityWarning / SeverityHint references in internal/audit/
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/audit.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_12
created_at: "2026-04-17T10:01:13Z"
completed_at: "2026-04-17T10:27:43Z"
---
