---
id: REQ-0060
kind: requirement
name: Plan Requires Approval Before Execution
slug: plan-requires-approval-before-execution-vldv
relationships:
    - target: plan-sk33
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:40Z"
statement: While a plan has plan_status draft, the syde CLI shall not allow any task under that plan to transition to in_progress or completed.
req_type: functional
priority: must
verification: integration test starting a task on a draft plan
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: Approval gates prevent implementation work from starting on unreviewed designs.
---
