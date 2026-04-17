---
id: REQ-0053
kind: requirement
name: Plan Phase ID Unique Within Plan
slug: plan-phase-id-unique-within-plan-1wq7
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:01Z"
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
---
