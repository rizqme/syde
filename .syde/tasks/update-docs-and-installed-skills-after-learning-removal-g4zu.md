---
id: TSK-0116
kind: task
name: Update docs and installed skills after learning removal
slug: update-docs-and-installed-skills-after-learning-removal-g4zu
relationships:
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: remove-learning-entity-and-cap-requirement-fanout-wqyz
      type: belongs_to
updated_at: "2026-04-15T09:23:33Z"
task_status: completed
priority: medium
objective: Remove learning command/entity references from skill docs and refresh installed skill copies.
details: Update SKILL.md, codex skill, command/entity references, installer hints, then run make install and syde install-skill --all.
acceptance: Installed skill files no longer document learning entities or syde remember.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/codex/SKILL.md
    - skill/references/commands.md
    - skill/references/entity-spec.md
    - internal/skill/installer.go
plan_ref: remove-learning-entity-and-cap-requirement-fanout-wqyz
plan_phase: phase_4
created_at: "2026-04-15T08:23:05Z"
completed_at: "2026-04-15T09:15:50Z"
---
