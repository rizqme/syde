---
id: TSK-0107
kind: task
name: Extend AuditedOverlap struct with Distinction field
slug: extend-auditedoverlap-struct-with-distinction-field-ze4t
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: RequirementEntity.AuditedOverlaps is []AuditedOverlap{Slug, Distinction} with YAML round-trip for both new and legacy slug-only entries
details: Edit internal/model (entity.go or the requirement struct file) to introduce type AuditedOverlap struct; implement custom UnmarshalYAML that accepts either a string (treated as {slug, distinction=''}) or a map. MarshalYAML always writes the full map form. Update any code referencing AuditedOverlaps slice.
acceptance: go build succeeds; round-trip of a requirement with legacy slug-only entries preserves behavior; new entries write distinction
affected_entities:
    - entity-model-f28o
    - cli-commands-hpjb
    - audit-engine-4ktg
affected_files:
    - internal/model/entity.go
    - internal/cli/add.go
    - internal/cli/update.go
    - internal/cli/requirements.go
    - internal/audit/requirements.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_1
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:10:30Z"
---
