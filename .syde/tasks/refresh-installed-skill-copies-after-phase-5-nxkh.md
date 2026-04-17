---
id: TSK-0020
kind: task
name: Refresh installed skill copies after Phase 5
slug: refresh-installed-skill-copies-after-phase-5-nxkh
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: Installed Claude and Codex skill copies reflect the new plan workflow.
details: make install && syde install-skill --all. Verify rg 'plan add-change' .claude/skills .agents/skills returns matches.
acceptance: Installed skill files document the new plan workflow.
affected_entities:
    - skill-installer-wbmu
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_5
created_at: "2026-04-15T11:42:14Z"
completed_at: "2026-04-15T12:41:18Z"
---
