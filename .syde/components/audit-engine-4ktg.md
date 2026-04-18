---
id: COM-0019
kind: component
name: Audit Engine
slug: audit-engine-4ktg
description: Shared health-check engine backing syde validate, syde sync check, and syde files commands
purpose: Give every CLI health command a single source of truth for findings instead of re-implementing the same checks in three places
files:
- internal/audit/audit.go
- internal/audit/cycles.go
- internal/audit/entity_fields.go
- internal/audit/orphans.go
- internal/audit/plan_phases.go
- internal/audit/relationships.go
- internal/audit/screens.go
- internal/audit/graph_rules.go
- internal/audit/plan_completion.go
- internal/audit/plan_authoring.go
- internal/audit/requirements.go
- internal/audit/surfaces.go
- internal/audit/surfaces_test.go
- internal/audit/symmetry.go
- internal/audit/symmetry_test.go
- internal/audit/severity_test.go
- internal/audit/overlap_test.go
- cmd/listoverlaps/main.go
- internal/audit/bidirectional.go
relationships:
- target: storage-engine
  type: depends_on
- target: entity-model
  type: depends_on
- target: summary-tree
  type: depends_on
- type: belongs_to
  target: syde-5tdt
- type: belongs_to
  target: syded-dashboard-e82c
updated_at: '2026-04-18T09:27:35Z'
responsibility: Produce Findings (a single strict severity level) covering entity field validation, relationship integrity, cycles, tree file references, orphan detection, file drift, contract surface coverage, flow coverage, and planning-post-plan symmetry
capabilities:
- Run full audit with per-category opt-out
- Detect orphan source files (non-ignored tree files with no owning entity)
- Detect file drift (file mtime newer than owner UpdatedAt)
- Enumerate file coverage (path → owner entities)
- Detect cyclic system nesting and component dependencies
boundaries: Does NOT mutate the store or the tree. Does NOT format output — callers handle printing. Does NOT call the CLI. Does NOT infer contracts or flows from free-form prose beyond the declared pattern extractors.
---
