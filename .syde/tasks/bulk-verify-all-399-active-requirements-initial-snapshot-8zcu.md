---
id: TSK-0245
kind: task
name: Bulk-verify all 399 active requirements (initial snapshot)
slug: bulk-verify-all-399-active-requirements-initial-snapshot-8zcu
updated_at: '2026-04-18T08:16:16Z'
task_status: completed
priority: high
objective: Every active req has a fully populated verified_against snapshot matching current code state
details: 'Shell loop: ''syde list requirement --format json | jq -r ... | xargs -I {} syde requirement verify {}''. Skip non-active reqs.'
acceptance: syde sync check exits 0 with the requirement_stale finding zero on a clean repo
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_3
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:16:16Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: syde-requirement-verify-shall-snapshot-sha-256-hashes-for-every-file-in-each-refining-component-3p42
---
