---
id: TSK-0240
kind: task
name: Implement requirement_no_component audit finding
slug: implement-requirementnocomponent-audit-finding-jn5m
updated_at: '2026-04-18T08:12:01Z'
task_status: completed
priority: high
objective: Active req with zero refines:component edges produces a finding under CatRequirement
acceptance: Removing refines:component from a test req causes 'syde sync check' to surface a finding mentioning that req's slug
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
  target: active-requirement-shall-refine-at-least-one-component-mke4
---
