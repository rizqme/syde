---
id: REQ-0055
kind: requirement
name: Plan Phase Auto-Completes On Final Task Done
slug: plan-phase-auto-completes-on-final-task-done-q9n6
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:31Z"
statement: While a plan phase is in_progress and all of its tasks have status completed, the syde CLI shall transition the phase to completed.
req_type: functional
priority: must
verification: integration test marking the last task done and asserting phase status
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Automatic completion reduces manual bookkeeping and keeps phase state consistent with task state.
---
