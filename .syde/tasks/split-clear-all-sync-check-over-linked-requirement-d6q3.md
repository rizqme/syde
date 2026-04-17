---
id: TSK-0104
kind: task
name: Split Clear-all-sync-check over-linked requirement
slug: split-clear-all-sync-check-over-linked-requirement-d6q3
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-remaining-sync-check-drift
      type: references
updated_at: "2026-04-17T10:46:19Z"
task_status: completed
objective: The Clear-all-sync-check approved-plan requirement is no longer over the 10-per-kind task cap
details: Enumerate the 20 linking tasks; author 2 child requirements with scoped names ('concept redesign tasks' / 'audit and overlap tasks'), each derives_from the parent; repoint each task via two-step --remove-rel + --add-rel <child>:references
acceptance: The parent requirement has <=10 inbound task links and each of the 2 child requirements has <=10 inbound task links
affected_entities:
    - clear-all-sync-check-concept-redesign-tasks
    - clear-all-sync-check-audit-and-overlap-tasks
plan_ref: clear-all-remaining-sync-check-drift-aokb
plan_phase: phase_4
created_at: "2026-04-17T08:48:21Z"
completed_at: "2026-04-17T09:11:19Z"
---
