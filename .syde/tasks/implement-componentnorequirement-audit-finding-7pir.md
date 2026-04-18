---
id: TSK-0241
kind: task
name: Implement component_no_requirement audit finding
slug: implement-componentnorequirement-audit-finding-7pir
updated_at: '2026-04-18T08:12:01Z'
task_status: completed
priority: high
objective: Component with files mapped and zero incoming refines from active reqs produces a finding under CatTraceability
acceptance: Removing all refining reqs from a component-with-files causes 'syde sync check' to surface a finding mentioning that component's slug
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
  target: component-with-mapped-files-shall-have-at-least-one-refining-requirement-300f
---
