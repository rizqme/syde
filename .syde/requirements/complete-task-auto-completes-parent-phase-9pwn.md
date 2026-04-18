---
id: REQ-0262
kind: requirement
name: Complete Task Auto-Completes Parent Phase
slug: complete-task-auto-completes-parent-phase-9pwn
relationships:
    - target: complete-task-k8je
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:05Z"
statement: When syde task done finishes the last task of a phase, the syde CLI shall auto-complete the parent phase and report phase_auto_completed as true.
req_type: interface
priority: must
verification: integration test completing the final task of a phase
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Automatic phase rollup keeps plan progress consistent without manual bookkeeping.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:05Z"
---
