---
id: TSK-0128
kind: task
name: Update skill docs and refresh installed copies
slug: update-skill-docs-and-refresh-installed-copies-vlro
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:44:33Z"
task_status: completed
priority: medium
objective: SKILL.md, entity-spec.md, commands.md cover the new requirement model and EARS conventions; installed copies match.
details: Edit skill/SKILL.md, skill/references/entity-spec.md, skill/references/commands.md, skill/codex/SKILL.md. Run make install and syde install-skill --all.
acceptance: Installed copies document req_type/priority/verification/refines/EARS; rg'd against the installed trees matches the source.
affected_entities:
    - skill-installer-wbmu
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:44:33Z"
---
