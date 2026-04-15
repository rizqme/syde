---
id: REQ-0069
kind: requirement
name: Plan Phase Completion Gate
slug: plan-phase-completion-gate-1jat
relationships:
    - target: plan-sk33
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:42Z"
statement: If any task under a plan phase is not completed, then the syde CLI shall not allow the phase to transition to completed.
req_type: constraint
priority: must
verification: integration test completing a phase with an outstanding task
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: Prevents phases from being marked done while work remains.
---
