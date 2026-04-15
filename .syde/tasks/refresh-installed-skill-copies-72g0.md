---
id: TSK-0104
kind: task
name: Refresh installed skill copies
slug: refresh-installed-skill-copies-72g0
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:08:42Z"
task_status: completed
priority: medium
objective: Reinstall the updated syde skills so live Claude/Codex skill files match the templates.
details: Run syde install-skill --all after source skill edits, then verify installed skill files contain the updated clarification and task-sync guidance and hooks still use syde command paths.
acceptance: Installed skill files contain the updated clarification and task-sync guidance.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/codex/SKILL.md
    - skill/references/commands.md
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_1
created_at: "2026-04-15T06:55:19Z"
completed_at: "2026-04-15T07:08:42Z"
---
