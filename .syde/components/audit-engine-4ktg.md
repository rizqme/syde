---
id: COM-0019
kind: component
name: Audit Engine
slug: audit-engine-4ktg
description: Shared health-check engine backing syde validate, syde sync check, and syde files commands
purpose: Give every CLI health command a single source of truth for findings instead of re-implementing the same checks in three places
files:
    - internal/audit/audit.go
    - internal/audit/concepts.go
    - internal/audit/cycles.go
    - internal/audit/entity_fields.go
    - internal/audit/orphans.go
    - internal/audit/plan_phases.go
    - internal/audit/relationships.go
    - internal/audit/screens.go
    - internal/audit/graph_rules.go
    - internal/audit/requirements.go
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: depends_on
    - target: entity-model
      type: depends_on
    - target: summary-tree
      type: depends_on
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
updated_at: "2026-04-15T11:06:07Z"
responsibility: Produce categorized Findings (errors/warnings/hints) covering entity field validation, relationship integrity, cycles, tree file references, orphan detection, and file drift
capabilities:
    - Run full audit with per-category opt-out
    - Detect orphan source files (non-ignored tree files with no owning entity)
    - Detect file drift (file mtime newer than owner UpdatedAt)
    - Enumerate file coverage (path → owner entities)
    - Detect cyclic system nesting and component dependencies
boundaries: Does NOT mutate the store or the tree. Does NOT format output — callers handle printing. Does NOT call the CLI.
---
