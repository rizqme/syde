---
id: REQ-0262
kind: requirement
name: Complete Task Auto-Completes Parent Phase
slug: complete-task-auto-completes-parent-phase-9pwn
relationships:
    - target: complete-task-k8je
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When syde task done finishes the last task of a phase, the syde CLI shall auto-complete the parent phase and report phase_auto_completed as true.
req_type: interface
priority: must
verification: integration test completing the final task of a phase
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Automatic phase rollup keeps plan progress consistent without manual bookkeeping.
---
