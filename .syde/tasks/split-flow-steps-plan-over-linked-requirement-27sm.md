---
id: TSK-0105
kind: task
name: Split Flow-steps plan over-linked requirement
slug: split-flow-steps-plan-over-linked-requirement-27sm
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-remaining-sync-check-drift
      type: references
updated_at: "2026-04-17T10:46:19Z"
task_status: completed
objective: The Flow-steps approved-plan requirement is no longer over the 10-per-kind cap for flows or tasks
details: Enumerate 30 linking flows + 19 linking tasks; author 3 flow-scoped children and 2 task-scoped children; repoint every linking entity onto its assigned child via the two-step rel edit pattern.
acceptance: The parent requirement has <=10 inbound flow links and <=10 inbound task links, all 5 children have <=10 inbound links of their respective kind
affected_entities:
    - flow-steps-plan-lifecycle-flows
    - flow-steps-entity-operation-flows
    - flow-steps-dashboard-browsing-flows
    - flow-steps-flow-authoring-tasks
    - flow-steps-chart-and-doc-tasks
plan_ref: clear-all-remaining-sync-check-drift-aokb
plan_phase: phase_4
created_at: "2026-04-17T08:48:21Z"
completed_at: "2026-04-17T09:15:19Z"
---
