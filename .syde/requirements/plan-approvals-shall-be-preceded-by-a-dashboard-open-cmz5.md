---
id: REQ-0337
kind: requirement
name: Plan approvals shall be preceded by a dashboard open
slug: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
description: Plan approval handoff requirement for opening the dashboard plan detail.
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:11Z"
statement: When the syde planning workflow drafts a plan and is ready to ask for user approval, the agent shall run syde plan open <plan> before presenting the approval prompt.
req_type: usability
priority: must
verification: 'Manual workflow inspection: every approval request in a session is preceded by a syde plan open invocation.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Manual URL navigation breaks reviewer flow. The agent should drive the dashboard tab directly.
audited_overlaps:
    - slug: plans-shall-pass-syde-plan-check-before-approval-0jkc
      distinction: Requires running syde plan open to surface the dashboard; target requires running syde plan check and fixing ERRORs.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:11Z"
---
