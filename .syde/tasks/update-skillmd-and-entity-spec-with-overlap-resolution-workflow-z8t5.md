---
id: TSK-0114
kind: task
name: Update SKILL.md and entity-spec with overlap resolution workflow
slug: update-skillmd-and-entity-spec-with-overlap-resolution-workflow-z8t5
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Skill docs teach MERGE/RENAME/DISTINCT and the slug:reason syntax on --audited
details: Add a 'Requirement overlap resolution' subsection to skill/SKILL.md under Requirement rules. Document the three outcomes, when each applies, and the slug:reason syntax. Update skill/references/entity-spec.md's audited_overlaps row. Reinstall via syde install-skill --all.
acceptance: skill/SKILL.md contains the new section; syde install-skill --all writes the updated files into .claude/skills and .agents/skills
affected_entities:
    - skill-7fmf
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_6
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:30:16Z"
---
