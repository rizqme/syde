---
acceptance: curl '/api/<proj>/query?mode=search&q=storage&kind=component' returns non-empty hits with snippets; curl '/api/<proj>/query?mode=by-file&path=internal/storage/index.go' returns owners + related.
affected_entities:
    - http-api
    - query-engine
affected_files:
    - internal/dashboard/api_readall.go
    - internal/dashboard/api.go
completed_at: "2026-04-14T08:19:48Z"
created_at: "2026-04-14T08:07:48Z"
details: 'In api_readall.go handleQueryAPI switch: case ''search'' parses kind/tag/limit/any/q and calls eng.Search(opts); case ''by-file'' parses path/with_related and calls eng.ByFile. Reuse writeSummaries for list formats; add writeSearchHits for search snippets.'
id: TSK-0045
kind: task
name: 'Dashboard handleQueryAPI: add search + by-file modes'
objective: Both new modes reachable via /api/<proj>/query with format=rich|json|compact|refs
plan_phase: phase_3
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dashboard-handlequeryapi-add-search-by-file-modes-epvv
task_status: completed
updated_at: "2026-04-14T08:19:48Z"
---
