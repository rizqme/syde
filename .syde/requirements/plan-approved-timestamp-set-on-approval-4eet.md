---
id: REQ-0073
kind: requirement
name: Plan Approved Timestamp Set On Approval
slug: plan-approved-timestamp-set-on-approval-4eet
relationships:
    - target: plan-sk33
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:46Z"
statement: When a plan transitions from draft to approved, the syde CLI shall set approved_at to the current wall clock time.
req_type: functional
priority: must
verification: integration test approving a plan and asserting approved_at is populated
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: Timestamps provide audit evidence of when work became eligible to start.
---
