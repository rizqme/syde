---
acceptance: Create a throwaway plan with an empty-ID phase (via direct yaml hand-edit for the test), run syde sync check, see an ERROR with the plan slug + phase index. sync check --errors-only exits non-zero.
affected_entities:
    - audit-engine
affected_files:
    - internal/audit/audit.go
completed_at: "2026-04-14T08:59:42Z"
created_at: "2026-04-14T08:55:51Z"
details: 'internal/audit: add CatPlanPhase constant. New file plan_phases.go (or append to audit.go) with planPhasesFindings(all) that walks every plan entity, builds per-plan idSet and parent map, emits Finding entries for each defect, severity ERROR. Integrate into audit.Run() alongside existing checks. DFS with on-stack set for cycle detection, parameterized per-plan.'
id: TSK-0050
kind: task
name: 'Audit: ERROR findings for corrupt plan phase data'
objective: syde sync check --strict reports empty-ID phases, duplicate IDs, self-parent, parent cycles, and dangling parent refs so corrupt plans block session end
plan_phase: phase_3
plan_ref: guard-plan-phase-integrity
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: audit-error-findings-for-corrupt-plan-phase-data-47lf
task_status: completed
updated_at: "2026-04-14T08:59:42Z"
---
