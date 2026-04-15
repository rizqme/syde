---
acceptance: syde query --code ConceptEntity returns at least the line in internal/model/entity.go where the struct is defined, plus the owning Entity Model component.
affected_entities:
    - query-engine
    - audit-engine
affected_files:
    - internal/query/engine.go
completed_at: "2026-04-14T09:22:27Z"
created_at: "2026-04-14T09:20:18Z"
details: 'internal/query/engine.go: SearchCodeOptions {Pattern, Limit, OwnerKind, OwnerSlug}. Engine.SearchCode walks the summary tree (audit.FileCoverage already loads it; expose a thin wrapper) for tracked, non-ignored files. exec.LookPath(''rg''); if found, run rg with literal -F mode and parse path:line:content. If rg missing, walk files in Go and Contains-check per line. Each hit becomes CodeHit{Path, Line, Snippet, OwnerKind, OwnerSlug, OwnerName}. Limit caps the result at 50 by default. Errors from rg should fall back to the Go walker silently.'
id: TSK-0051
kind: task
name: Engine.SearchCode with rg fallback to Go walker
objective: syde query --code 'pattern' returns ranked source-file hits with owning entity attached
plan_phase: phase_1
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: enginesearchcode-with-rg-fallback-to-go-walker-o098
task_status: completed
updated_at: "2026-04-14T09:22:27Z"
---
