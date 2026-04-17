---
id: TSK-0040
kind: task
name: Refresh installed skill copies after Phase 6
slug: refresh-installed-skill-copies-after-phase-6-nu5v
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: Installed Claude and Codex skill copies teach the new syde plan check step.
details: make install && syde install-skill --all. rg syde plan check .claude/skills .agents/skills returns matches.
acceptance: Installed skill files document the new gate.
affected_entities:
    - skill-installer-wbmu
    - web-spa-jy9z
affected_files:
    - skill/hooks.json
    - skill/codex/hooks.json
    - skill/SKILL.md
    - skill/codex/SKILL.md
    - web/src/components/PlanChangesView.tsx
    - internal/skill/installer.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_6
created_at: "2026-04-15T13:15:56Z"
completed_at: "2026-04-16T01:00:13Z"
---
