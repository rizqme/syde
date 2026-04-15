---
acceptance: Full-text search returns SearchHits with non-empty id/kind/name/slug. Unit-smoke via ./syde query --search storage.
affected_entities:
    - storage-engine
    - query-engine
affected_files:
    - internal/storage/index.go
    - internal/storage/indexer.go
completed_at: "2026-04-14T08:13:48Z"
created_at: "2026-04-14T08:07:21Z"
details: 'IndexWords now marshals JSON {file_ref, field, word} into w: values. Search returns []SearchHit with kind, id, name, slug, file, field, word populated.'
id: TSK-0041
kind: task
name: 'Reshape w: index values to carry full FileRef + field + word'
objective: Every inverted-index hit resolves back to a complete entity identity without a second lookup
plan_phase: phase_1
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: reshape-w-index-values-to-carry-full-fileref-field-word-5uo1
task_status: completed
updated_at: "2026-04-14T08:13:48Z"
---
