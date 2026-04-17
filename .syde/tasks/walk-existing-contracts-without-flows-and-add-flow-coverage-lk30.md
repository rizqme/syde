---
id: TSK-0124
kind: task
name: Walk existing contracts without flows and add flow coverage
slug: walk-existing-contracts-without-flows-and-add-flow-coverage-lk30
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every active contract participates in at least one flow step; sync check reports zero contract_flow_coverage errors
details: Run syde sync check, collect contractFlowFindings; for each missing contract, either extend an existing flow with a new step that references it via --step, or author a new flow for contracts that warrant their own journey. Contract coverage plan (phases 6-9) may have added new contracts that need matching flows — do them together.
acceptance: syde sync check reports zero contractFlowFindings errors; every active contract has at least one step reference
affected_entities:
    - review-plan-c4qb
    - execute-plan-syvb
    - add-entity-9m4a
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_11
created_at: "2026-04-17T09:50:47Z"
completed_at: "2026-04-17T10:43:19Z"
---
