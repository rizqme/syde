---
id: REQ-0379
kind: requirement
name: 'Approved plan: Clear all remaining sync check drift'
slug: approved-plan-clear-all-remaining-sync-check-drift-phkm
relationships:
    - target: clear-all-remaining-sync-check-drift-aokb
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:19:28Z"
statement: The syde design model shall reconcile every residual sync-check finding by mapping orphan files to components, authoring plan-declared requirements, executing plan-declared deletions, and splitting any approved-plan requirement that exceeds the 10-per-kind link cap.
req_type: constraint
priority: must
verification: syde sync check --strict exits 0 and every older plan can syde plan complete without --force
source: plan
source_ref: plan:clear-all-remaining-sync-check-drift-aokb
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T09:04:08Z"
---
