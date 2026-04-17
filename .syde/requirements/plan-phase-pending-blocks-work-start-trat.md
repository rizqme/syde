---
id: REQ-0059
kind: requirement
name: Plan Phase Pending Blocks Work Start
slug: plan-phase-pending-blocks-work-start-trat
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:01Z"
statement: While a plan phase has status pending, the syde CLI shall not allow any of its tasks to transition to in_progress without first starting the phase.
req_type: functional
priority: should
verification: integration test starting a task on a pending phase
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Enforces the pending to in_progress ordering so work does not begin prematurely.
audited_overlaps:
    - slug: task-pending-cannot-be-completed-directly-f72j
      distinction: Blocks task start while phase is pending; target blocks task jumping from pending directly to completed without in_progress.
---
