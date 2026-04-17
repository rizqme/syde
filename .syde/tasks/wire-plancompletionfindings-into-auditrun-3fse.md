---
id: TSK-0010
kind: task
name: Wire planCompletionFindings into audit.Run
slug: wire-plancompletionfindings-into-auditrun-3fse
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: planCompletionFindings runs as part of the default audit pipeline so syde sync check --strict enforces plan integrity.
details: internal/audit/audit.go Run adds rep.Findings = append(rep.Findings, planCompletionFindings(all)...) under the existing Relationship section or a new PlanIntegrity section. Add a new FindingCategory constant if needed.
acceptance: syde sync check --strict surfaces plan completion findings when present.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/audit.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_3
created_at: "2026-04-15T11:41:14Z"
completed_at: "2026-04-15T12:01:40Z"
---
