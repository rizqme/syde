---
acceptance: syde query --file internal/model/plan.go --content shows owner (Entity Model) + related entities + the source of plan.go in the same output.
affected_entities:
    - query-engine
affected_files:
    - internal/query/engine.go
    - internal/query/formatter.go
completed_at: "2026-04-14T09:23:36Z"
created_at: "2026-04-14T09:20:18Z"
details: 'internal/query/engine.go: ByFile signature gains withContent bool (or extend SearchOptions-like ByFileOptions for clarity). ByFileResult gains Content string + ContentBytes int + ContentTruncated bool. Read the file when exact match, cap at 100KB, mark truncated. Skip when prefix mode (multiple files don''t make sense to inline). internal/query/formatter.go FormatByFile prints content under owners + related when present.'
id: TSK-0052
kind: task
name: ByFile --content inlines file body with size cap
objective: syde query --file <path> --content returns owners + related + the file content in one call
plan_phase: phase_1
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: byfile-content-inlines-file-body-with-size-cap-up7n
task_status: completed
updated_at: "2026-04-14T09:23:36Z"
---
