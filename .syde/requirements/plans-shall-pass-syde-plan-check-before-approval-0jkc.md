---
id: REQ-0336
kind: requirement
name: Plans shall pass syde plan check before approval
slug: plans-shall-pass-syde-plan-check-before-approval-0jkc
description: Plan approval gate requiring a clean syde plan check.
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:07Z"
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
---
