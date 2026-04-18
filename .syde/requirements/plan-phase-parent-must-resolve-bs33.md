---
id: REQ-0051
kind: requirement
name: Plan Phase Parent Must Resolve
slug: plan-phase-parent-must-resolve-bs33
relationships:
    - target: plan-phase-23bb
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:37Z"
statement: If a plan phase declares a parent_phase that does not resolve to another phase in the same plan, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test adding a phase with an unknown parent_phase ID
source: manual
source_ref: concept:plan-phase-23bb
requirement_status: active
rationale: Dangling parent links corrupt the phase hierarchy.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:37Z"
---
