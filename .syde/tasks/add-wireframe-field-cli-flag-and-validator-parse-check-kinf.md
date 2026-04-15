---
acceptance: syde add contract Smoke --contract-kind screen --interaction-pattern render --input '/smoke' --input-parameter 'slug|string|route param' --output 'rendered UI' --output-parameter 'body|html|rendered markup' --wireframe '<screen><layout direction="vertical"><heading>Smoke</heading></layout></screen>' --add-rel 'syde:belongs_to' creates the contract; malformed UIML fails validation.
affected_entities:
    - entity-model
    - cli-commands
affected_files:
    - internal/model/entity.go
    - internal/model/validation.go
    - internal/cli/add.go
    - internal/cli/update.go
completed_at: "2026-04-14T11:30:00Z"
created_at: "2026-04-14T11:23:14Z"
details: 'internal/model/entity.go: add Wireframe string field to ContractEntity yaml:''wireframe,omitempty''. internal/model/validation.go case *ContractEntity: if v.ContractKind == "screen" require Wireframe non-empty and run uiml.Parse on it (check internal/uiml for cycle with model — if cycle exists, move parse check to internal/audit/screens.go instead). internal/cli/add.go + update.go register addContractWireframe / updContractWireframe string vars and --wireframe flag; wire into ContractEntity switch case.'
id: TSK-0076
kind: task
name: Add Wireframe field, CLI flag, and validator parse check
objective: ContractEntity.Wireframe persists; CLI accepts --wireframe; validator rejects malformed UIML when contract_kind=screen
plan_phase: phase_1
plan_ref: screen-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-wireframe-field-cli-flag-and-validator-parse-check-kinf
task_status: completed
updated_at: "2026-04-14T11:30:00Z"
---
