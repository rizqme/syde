---
acceptance: syde plan add-phase <slug> (no --name) errors cleanly without touching the file. syde plan add-phase --name 'x' still succeeds and the written phase has a non-empty id.
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/plan.go
completed_at: "2026-04-14T08:58:12Z"
created_at: "2026-04-14T08:55:51Z"
details: 'In internal/cli/plan.go planAddPhaseCmd: check strings.TrimSpace(addPhaseName) != '''' — error ''phase --name is required'' if blank. After computing phaseID, assert phaseID != '''' and that no existing p.Phases entry shares the same ID. Return fmt.Errorf with a one-line remediation hint on collision.'
id: TSK-0049
kind: task
name: Require --name and non-empty ID in syde plan add-phase
objective: Malformed phases cannot be created via the add-phase CLI
plan_phase: phase_2
plan_ref: guard-plan-phase-integrity
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: require-name-and-non-empty-id-in-syde-plan-add-phase-s253
task_status: completed
updated_at: "2026-04-14T08:58:12Z"
---
