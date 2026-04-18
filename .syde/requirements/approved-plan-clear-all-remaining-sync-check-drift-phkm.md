---
id: REQ-0379
kind: requirement
name: 'Approved plan: Clear all remaining sync check drift'
slug: approved-plan-clear-all-remaining-sync-check-drift-phkm
relationships:
    - target: clear-all-remaining-sync-check-drift-aokb
      type: references
      label: approved_plan
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:14Z"
statement: The syde design model shall reconcile every residual sync-check finding by mapping orphan files to components, authoring plan-declared requirements, executing plan-declared deletions, and splitting any approved-plan requirement that exceeds the 10-per-kind link cap.
req_type: constraint
priority: must
verification: syde sync check --strict exits 0 and every older plan can syde plan complete without --force
source: plan
source_ref: plan:clear-all-remaining-sync-check-drift-aokb
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T09:04:08Z"
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:14Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:14Z"
---
