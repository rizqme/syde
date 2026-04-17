---
id: TSK-0022
kind: task
name: Refresh summary tree after the revamp
slug: refresh-summary-tree-after-the-revamp-aou7
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: syde tree status --strict exits 0 after all Phase 1-5 edits land.
details: syde tree scan, dispatch subagents to summarize stale leaves per the standard leaves-first loop, summarize stale folders, verify exit 0.
acceptance: syde tree status --strict returns exit 0.
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_6
created_at: "2026-04-15T11:42:14Z"
completed_at: "2026-04-15T12:46:06Z"
---
