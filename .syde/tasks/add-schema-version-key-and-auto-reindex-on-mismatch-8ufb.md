---
acceptance: Delete .syde/.index/, start syded, index rebuilds without manual reindex. meta:schema key present after startup.
affected_entities:
    - storage-engine
affected_files:
    - internal/storage/index.go
    - internal/storage/store.go
completed_at: "2026-04-14T08:12:11Z"
created_at: "2026-04-14T08:07:21Z"
details: Add const IndexSchemaVersion=2 in index.go. In storage.Store.Open (or wherever Index opens), read 'meta:schema' key; if missing or <current, call Reindex and write new version. No-op on subsequent starts.
id: TSK-0040
kind: task
name: Add schema version key and auto reindex on mismatch
objective: syded rebuilds index automatically when schema version is missing or older than current
plan_phase: phase_1
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-schema-version-key-and-auto-reindex-on-mismatch-8ufb
task_status: completed
updated_at: "2026-04-14T08:12:11Z"
---
