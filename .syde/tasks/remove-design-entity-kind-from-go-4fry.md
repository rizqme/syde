---
id: TSK-0121
kind: task
name: Remove design entity kind from Go
slug: remove-design-entity-kind-from-go-4fry
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:28:00Z"
task_status: completed
priority: high
objective: KindDesign, DesignEntity, CLI design subcommand, dashboard design view, and skill docs are gone.
details: Same removal pattern. Zero design entities exist on disk so no data migration needed. Remove internal/model/design.go, internal/cli/design.go, any dashboard design handlers, skill docs references. Drop .syde/designs directory.
acceptance: rg KindDesign internal returns zero matches; go build ./... succeeds.
affected_entities:
    - entity-model-f28o
    - http-api-afos
    - cli-commands-hpjb
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_1
created_at: "2026-04-15T09:53:21Z"
completed_at: "2026-04-15T10:28:00Z"
---
