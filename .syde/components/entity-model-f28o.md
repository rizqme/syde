---
id: COM-0003
kind: component
name: Entity Model
slug: entity-model-f28o
description: Typed entity schema, kind dispatch, and per-kind validation rules.
purpose: Define the typed schema for every kind of syde entity
notes:
    - validation.go now requires description on every entity kind (2026-04-14).
files:
    - internal/model/entity.go
    - internal/model/relationship.go
    - internal/model/plan.go
    - internal/model/plan_test.go
    - internal/model/task.go
    - internal/model/validation.go
relationships:
    - target: syde-cli
      type: belongs_to
updated_at: "2026-04-17T10:52:48Z"
responsibility: BaseEntity + per-kind structs + validation rules + plan/task/relationship types + AuditedOverlap with distinction rationale
capabilities:
    - Type-switched entity constructors (NewEntityForKind)
    - YAML (un)marshal with frontmatter-compatible tags
    - Per-kind validation rules (required fields, acyclic deps, target existence)
    - ID prefix mapping (SYS, COM, CON, CPT, FLW, DEC, PLN, TSK, DSG, LRN)
boundaries: Does NOT persist entities (storage owns that). Does NOT resolve relationships (query/graph own that).
---
