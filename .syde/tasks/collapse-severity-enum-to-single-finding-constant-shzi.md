---
id: TSK-0125
kind: task
name: Collapse Severity enum to single Finding constant
slug: collapse-severity-enum-to-single-finding-constant-shzi
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-review-strict-severity-verify-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: internal/audit defines one Severity constant (Finding); SeverityError, SeverityWarning, SeverityHint removed
details: 'Edit internal/audit/audit.go: replace the three Severity constants with a single Finding. Update Report.Counts and any formatting that switched on severity to report only Finding count. Ensure the JSON health envelope still round-trips.'
acceptance: go build succeeds; grep 'SeverityWarning\|SeverityHint' returns zero results
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/audit.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_12
created_at: "2026-04-17T10:01:13Z"
completed_at: "2026-04-17T10:27:43Z"
---
