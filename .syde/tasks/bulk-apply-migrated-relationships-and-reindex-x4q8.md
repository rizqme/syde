---
id: TSK-0237
kind: task
name: Bulk-apply migrated relationships and reindex
slug: bulk-apply-migrated-relationships-and-reindex-x4q8
updated_at: '2026-04-18T08:08:01Z'
task_status: completed
priority: high
objective: All 207 reqs have refines:component added (per worksheet) and belongs_to:system removed in a single batch via direct YAML edit + syde reindex
details: Python script reads each req YAML, removes belongs_to entries whose target resolves to kind=system, appends refines entries for assigned components. Writes back. Then 'syde reindex' rebuilds Badger.
acceptance: syde reindex completes without errors; spot-check 5 random reqs show refines:component present and belongs_to:system absent
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_1
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:08:01Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: active-requirement-shall-refine-at-least-one-component-mke4
---
