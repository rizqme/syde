---
id: TSK-0239
kind: task
name: Add VerifiedAgainst field to RequirementEntity
slug: add-verifiedagainst-field-to-requiremententity-mqxl
updated_at: '2026-04-18T08:09:18Z'
task_status: completed
priority: high
objective: RequirementEntity carries a yaml-serialized verified_against map keyed by component canonical slug
details: Add VerifiedAgainst map[string]VerifiedSnapshot field with yaml:verified_against,omitempty tag; define VerifiedSnapshot struct with Hash and At fields. Both omitempty so existing reqs serialize unchanged.
acceptance: go build succeeds; existing requirement YAML files round-trip without diff when verified_against is empty
affected_files:
- internal/model/entity.go
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_2
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:09:18Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: requirement-shall-be-marked-stale-when-refining-component-file-content-changes-85v0
---
