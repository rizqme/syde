---
acceptance: syde query --search 'relationship label' returns at least the Relationship concept hit, with a 'broadened (no exact match for all tokens)' note in the output.
affected_entities:
    - query-engine
affected_files:
    - internal/query/engine.go
    - internal/query/formatter.go
completed_at: "2026-04-14T09:30:00Z"
created_at: "2026-04-14T09:20:18Z"
details: 'internal/query/engine.go: in Engine.Search, after the AND pass yields filtered=[] and !opts.Any, retry the same pipeline with opts.Any=true, set hit.Broadened=true on every result. SearchHit gains Broadened bool json field. internal/query/formatter.go FormatSearchHits prepends a one-line note when any hit is Broadened.'
id: TSK-0055
kind: task
name: Engine.Search AND→OR fallback with Broadened flag
objective: Multi-word queries that yield zero AND hits automatically retry as OR and label the results as broadened so agents understand the relaxation
plan_phase: phase_2
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: enginesearch-andor-fallback-with-broadened-flag-kkhu
task_status: completed
updated_at: "2026-04-14T09:30:00Z"
---
