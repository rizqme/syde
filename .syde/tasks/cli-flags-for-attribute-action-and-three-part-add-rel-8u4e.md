---
acceptance: syde add concept X --meaning m --attribute 'a|string|x' --action 'go|does' --add-rel 'y:relates_to:one-to-many' writes a file whose on-disk YAML has attributes, actions, and a relates_to with label 'one-to-many'.
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/add.go
    - internal/cli/update.go
    - internal/cli/helpers.go
completed_at: "2026-04-14T09:49:03Z"
created_at: "2026-04-14T09:44:26Z"
details: 'internal/cli/add.go: addConceptAttributes and addConceptActions []string flag slices, registered with StringArrayVar. parseConceptAttributes() and parseConceptActions() helpers. Concept case wires v.Attributes = parseConceptAttributes(addConceptAttributes), same for Actions. internal/cli/update.go mirrors the same with cmd.Flags().Changed gating. Extract shared --add-rel parser into a helper (internal/cli/helpers.go or a new file) that accepts 2- or 3-part specs; add.go and update.go both call it.'
id: TSK-0058
kind: task
name: CLI flags for --attribute, --action, and three-part --add-rel
objective: syde add concept / syde update concept accept the new flags and --add-rel optionally carries a cardinality label
plan_phase: phase_2
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: cli-flags-for-attribute-action-and-three-part-add-rel-8u4e
task_status: completed
updated_at: "2026-04-14T09:49:03Z"
---
