---
id: TSK-0127
kind: task
name: Write requirement derivation algorithm doc
slug: write-requirement-derivation-algorithm-doc-9gy7
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:40:38Z"
task_status: completed
priority: high
objective: skill/references/requirement-derivation.md exists as a deterministic per-kind procedure that a subagent can follow to produce EARS requirements from an entity.
details: Document the algorithm for Component, Contract, System, Concept (inputs, per-field procedures, EARS pattern selection, req_type inference, priority defaults, verification inference, back-link conventions, quality checks). Include worked examples.
acceptance: Doc is written and embedded via go:embed in internal/skill/templates.go, linked from SKILL.md and entity-spec.md.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/references/requirement-derivation.md
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:40:38Z"
---
