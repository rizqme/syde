---
id: REQ-0386
kind: requirement
name: Plan complete shall require clean sync check
slug: plan-complete-shall-require-clean-sync-check-9jcs
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:50Z"
statement: If the syde sync check reports any errors, then the syde plan complete command shall refuse to mark the plan completed without --force.
req_type: constraint
priority: must
verification: syde plan complete exits non-zero when sync check has errors unless --force is passed.
source: plan
requirement_status: active
rationale: Plan completion is the canonical gate; forcing should be rare and intentional.
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: plan-complete rule is a behavioural gate while the approved-plan requirement captures plan intent at approval time — distinct roles in the plan lifecycle
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:50Z"
---
