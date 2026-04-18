---
id: TSK-0242
kind: task
name: Implement requirement_targets_system audit finding
slug: implement-requirementtargetssystem-audit-finding-fb0p
updated_at: '2026-04-18T08:12:01Z'
task_status: completed
priority: high
objective: Active req with refines:system or belongs_to:system produces a finding
acceptance: Adding belongs_to:syde to a test req causes 'syde sync check' to surface a finding citing the offending edge
affected_files:
- internal/audit/requirements.go
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_2
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:12:01Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: requirement-shall-not-refine-or-belong-to-a-system-m4c2
---
