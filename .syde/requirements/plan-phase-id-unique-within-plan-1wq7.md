---
id: REQ-0053
kind: requirement
name: Plan Phase ID Unique Within Plan
slug: plan-phase-id-unique-within-plan-1wq7
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:51Z"
statement: The syde CLI shall ensure that every plan phase ID is unique within its parent plan.
req_type: constraint
priority: must
verification: integration test attempting to add two phases with the same phase_N id
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Non-unique phase IDs break task-to-phase lookups.
audited_overlaps:
    - slug: tree-node-path-unique-in-tree-zgei
      distinction: Enforces uniqueness of plan phase IDs within a plan; target enforces uniqueness of tree node paths in the summary tree.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:51Z"
---
