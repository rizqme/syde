---
id: TSK-0119
kind: task
name: Sync check errors on active requirements with uncovered contract surfaces
slug: sync-check-errors-on-active-requirements-with-uncovered-contract-surfaces-bycs
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: syde sync check emits ERROR for every active requirement whose statement surface is not covered by any active contract
details: Add a new finding generator (requirements.go or a new surfaces.go in audit) that walks active requirements, extracts surfaces, and compares against the set of active contract invocation signatures; emit ERROR for each gap, reporting the missing surface and suggesting either authoring a contract or rewording the requirement.
acceptance: syde sync check reports the current uncovered-surface gaps; once Claude fills them, sync check exits 0
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/requirements.go
    - internal/audit/audit.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_9
created_at: "2026-04-17T09:46:36Z"
completed_at: "2026-04-17T10:23:22Z"
---
