---
id: TSK-0100
kind: task
name: Author 9 Flow-steps plan requirements
slug: author-9-flow-steps-plan-requirements-ufhn
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-remaining-sync-check-drift
      type: references
updated_at: "2026-04-17T10:46:19Z"
task_status: completed
objective: Flow-steps plan's new-requirement declarations reference existing requirement names instead of the out-of-date draft names
details: 'For each of the 9 mismatched declarations in plan flow-steps-with-contract-references-and-flowchart-rendering-m0b5, remove the stale entry and add a fresh entry pointing at the actual implemented requirement name: Flow step shall have six fields, Each flow shall represent one user goal, Steps without on-success shall connect to next, Audit shall error on duplicate step IDs, Audit shall warn on steps with empty contract, Dashboard shall render flow steps as flowchart, Flows shall be tagged by category, Plans shall decompose into granular requirements, Dashboard shall render plan prose as markdown.'
acceptance: Flow-steps plan's show-changes shows the 9 new-req entries referencing names that resolve to existing requirement entities. Sync check reports no 'claims to create new requirement' errors for this plan.
affected_entities:
    - flow-steps-with-contract-references-and-flowchart-rendering-m0b5
plan_ref: clear-all-remaining-sync-check-drift-aokb
plan_phase: phase_2
created_at: "2026-04-17T08:48:21Z"
completed_at: "2026-04-17T09:06:19Z"
---
