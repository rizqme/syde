---
id: TSK-0122
kind: task
name: Refresh installed skill copies after Phase 1
slug: refresh-installed-skill-copies-after-phase-1-o512
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:31:23Z"
task_status: completed
priority: medium
objective: Installed Claude and Codex skill copies no longer reference decisions or designs.
details: make install, syde install-skill --all, verify installed copies are clean.
acceptance: rg decision .claude/skills .agents/skills returns zero matches for entity-kind references.
affected_entities:
    - skill-installer-wbmu
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_1
created_at: "2026-04-15T09:53:21Z"
completed_at: "2026-04-15T10:31:23Z"
---
