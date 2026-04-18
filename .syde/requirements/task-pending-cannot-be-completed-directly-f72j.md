---
id: REQ-0122
kind: requirement
name: Task Pending Cannot Be Completed Directly
slug: task-pending-cannot-be-completed-directly-f72j
relationships:
    - target: task-d3oc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:18Z"
statement: While a task has status pending, the syde CLI shall not allow it to transition directly to completed without first entering in_progress.
req_type: functional
priority: should
verification: integration test running syde task done on a pending task
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Enforces the pending to in_progress to completed ordering so work is tracked.
audited_overlaps:
    - slug: plan-phase-pending-blocks-work-start-trat
      distinction: Task status transition rule blocks pending to completed jumps; phase status rule blocks task work until the parent phase leaves pending.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:18Z"
---
