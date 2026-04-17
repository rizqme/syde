---
id: TSK-0007
kind: task
name: syde plan add-change new subcommand
slug: syde-plan-add-change-new-subcommand-shpt
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan add-change <plan> <kind> new --name --what --why plus kind-specific draft flags appends a NewChange.
details: Reuse existing add.go flag variables (addResponsibility, addCapabilities, addContractKind, addInput, addStatement, addReqType, etc.) by factoring a helper that returns a map[string]interface{} of populated fields for the requested kind. Validate that kind-required fields are present (component.responsibility, contract.input/output, requirement.statement matching EARS, etc.).
acceptance: Running add-change new contract with full screen flags writes a draft with contract_kind and wireframe populated.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_2
created_at: "2026-04-15T11:40:57Z"
completed_at: "2026-04-15T11:57:25Z"
---
