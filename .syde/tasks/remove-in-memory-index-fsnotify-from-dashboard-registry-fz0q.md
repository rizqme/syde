---
acceptance: go build passes; dashboard reads go through persistent badger; no dir-lock errors because syded is the only opener
completed_at: "2026-04-14T07:00:15Z"
details: Revert GetStore to call storage.NewStore (persistent). Delete startStoreWatcher. Remove fsnotify import if unused elsewhere.
id: TSK-0027
kind: task
name: Remove in-memory index + fsnotify from dashboard registry
objective: Dashboard opens persistent BadgerDB exclusively; fsnotify code deleted
plan_phase: phase_5
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: remove-in-memory-index-fsnotify-from-dashboard-registry-fz0q
task_status: completed
updated_at: "2026-04-14T07:00:15Z"
---
