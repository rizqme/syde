---
id: TSK-0037
kind: task
name: Wire planAuthoringFindings into audit.Run
slug: wire-planauthoringfindings-into-auditrun-zj85
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: audit.Run appends planAuthoringFindings results so syde sync check --strict picks them up.
details: Add the call alongside planCompletionFindings in audit.go Run.
acceptance: syde sync check --strict shows findings categorized as plan_authoring.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/audit.go
    - internal/audit/plan_authoring.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_6
created_at: "2026-04-15T13:15:56Z"
completed_at: "2026-04-15T21:38:25Z"
---
