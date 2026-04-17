---
id: TSK-0103
kind: task
name: Delete design-model-operations-coverage flow
slug: delete-design-model-operations-coverage-flow-r9kc
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-remaining-sync-check-drift
      type: references
updated_at: "2026-04-17T10:46:19Z"
task_status: completed
objective: The flow declared for deletion no longer exists on disk or in the index
details: syde remove design-model-operations-coverage --force; if it has inbound step refs from other flows, clean those first. Reindex after.
acceptance: syde query design-model-operations-coverage returns 404 and the Flow-steps plan's deletion declaration is satisfied
plan_ref: clear-all-remaining-sync-check-drift-aokb
plan_phase: phase_3
created_at: "2026-04-17T08:48:21Z"
completed_at: "2026-04-17T09:08:41Z"
---
