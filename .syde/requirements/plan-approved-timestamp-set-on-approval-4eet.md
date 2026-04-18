---
id: REQ-0073
kind: requirement
name: Plan Approved Timestamp Set On Approval
slug: plan-approved-timestamp-set-on-approval-4eet
relationships:
    - target: plan-sk33
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:10Z"
statement: When a plan transitions from draft to approved, the syde CLI shall set approved_at to the current wall clock time.
req_type: functional
priority: must
verification: integration test approving a plan and asserting approved_at is populated
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: Timestamps provide audit evidence of when work became eligible to start.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:10Z"
---
