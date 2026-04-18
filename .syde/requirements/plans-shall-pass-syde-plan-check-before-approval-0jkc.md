---
id: REQ-0336
kind: requirement
name: Plans shall pass syde plan check before approval
slug: plans-shall-pass-syde-plan-check-before-approval-0jkc
description: Plan approval gate requiring a clean syde plan check.
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: relates_to
updated_at: "2026-04-18T09:38:10Z"
statement: When the syde planning workflow drafts a plan, the agent shall run syde plan check on the plan and address every ERROR before presenting the plan for approval.
req_type: constraint
priority: must
verification: 'Manual workflow inspection: every approved plan in the repo must have been preceded by a syde plan check that exited 0.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Plan authoring gaps are the dominant failure mode. A programmatic gate catches them before approval instead of after execution drift.
audited_overlaps:
    - slug: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      distinction: plan check validates and addresses ERROR findings; plan open launches the dashboard view before approval, separate steps.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:10Z"
---
