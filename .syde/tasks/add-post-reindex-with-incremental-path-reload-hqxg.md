---
acceptance: POST with one path updates just that entity in index; subsequent GET returns fresh data
affected_entities:
    - http-api
affected_files:
    - internal/dashboard/api.go
completed_at: "2026-04-14T06:44:16Z"
created_at: "2026-04-14T06:38:16Z"
details: 'POST /api/<project>/reindex body {paths:[]string, full:bool}. For each path: FileStore.LoadOne + Store.indexEntity. full:true rebuilds the whole index via Reindex(). Returns {indexed:N, failed:[]}. Locking: single writer (syded) so no contention concerns.'
id: TSK-0016
kind: task
name: Add POST /reindex with incremental path reload
objective: syded reloads specific markdown paths into the index on CLI request — synchronous, returns when done
plan_phase: phase_1
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-post-reindex-with-incremental-path-reload-hqxg
task_status: completed
updated_at: "2026-04-14T06:44:16Z"
---
