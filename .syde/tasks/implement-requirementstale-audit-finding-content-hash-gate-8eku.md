---
id: TSK-0243
kind: task
name: Implement requirement_stale audit finding (content-hash gate)
slug: implement-requirementstale-audit-finding-content-hash-gate-8eku
updated_at: '2026-04-18T08:12:02Z'
task_status: completed
priority: high
objective: For each refines:component, if any file in component.files has SHA-256 differing from req.verified_against entry (or no entry exists), emit finding
details: Walk active reqs; for each refines:component edge, resolve component, hash each file in component.files, compare to req.VerifiedAgainst[componentSlug].Hash. Mismatch or missing entry → finding.
acceptance: After bulk-verifying all reqs in Phase 3, 'syde sync check' passes; touching any component file before re-verifying produces a stale finding for the refining req(s)
affected_files:
- internal/audit/requirements.go
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_2
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:12:02Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: requirement-shall-be-marked-stale-when-refining-component-file-content-changes-85v0
---
