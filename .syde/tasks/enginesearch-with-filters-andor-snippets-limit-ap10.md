---
acceptance: 'Go unit smoke or CLI smoke: ./syde query --search ''badger index'' --kind component --limit 3 returns at most 3 components with snippets.'
affected_entities:
    - query-engine
affected_files:
    - internal/query/engine.go
    - internal/storage/index.go
completed_at: "2026-04-14T08:17:26Z"
created_at: "2026-04-14T08:07:36Z"
details: 'SearchOptions struct {Query, Kind, Tag, Any bool, Limit int}. Pipeline: tokenize → per-token prefix scan (Index.SearchTokens returning map[token][]SearchHit) → intersect-by-entity-id (AND) or union (OR) → kind/tag filter → load description/body for snippet → truncate to Limit.'
id: TSK-0043
kind: task
name: Engine.Search with filters, AND/OR, snippets, limit
objective: Engine exposes a single Search(opts) entry point used by CLI and dashboard
plan_phase: phase_2
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: enginesearch-with-filters-andor-snippets-limit-ap10
task_status: completed
updated_at: "2026-04-14T08:17:26Z"
---
