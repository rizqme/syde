---
id: TSK-0120
kind: task
name: Remove decision entity kind from Go
slug: remove-decision-entity-kind-from-go-ps40
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:26:29Z"
task_status: completed
priority: high
objective: KindDecision, DecisionEntity, CLI flags, dashboard UI, API endpoints, and skill docs no longer mention decisions.
details: 'Same pattern as the learning removal: strip from internal/model/entity.go (enum + dispatch + KindPlural + IDPrefix + NewEntityForKind), delete internal/model/decision.go if standalone, remove decision-specific validation, query resolver cases, dashboard api.go + api_readall.go + html.go handlers, CLI add.go/update.go flags, skill references.'
acceptance: rg KindDecision internal returns zero matches; go build ./... succeeds.
affected_entities:
    - entity-model-f28o
    - audit-engine-4ktg
    - query-engine-9k84
    - http-api-afos
    - cli-commands-hpjb
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_1
created_at: "2026-04-15T09:53:21Z"
completed_at: "2026-04-15T10:26:29Z"
---
