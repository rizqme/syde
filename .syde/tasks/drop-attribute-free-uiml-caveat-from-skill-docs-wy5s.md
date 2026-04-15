---
acceptance: grep -ri 'attribute-free' skill/ returns zero hits. Reading the updated screen contract section produces a copy-pasteable example using <layout direction="horizontal">.
affected_entities:
    - skill-installer
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/references/commands.md
completed_at: "2026-04-15T03:11:21Z"
created_at: "2026-04-15T02:53:51Z"
details: 'skill/SKILL.md screen contract subsection: remove the ''stick to attribute-free structural tags'' bullet and the surrounding caveat paragraph. Replace with positive guidance pointing at the new wireframe vocabulary (sidebar/main/panel/card/grid/layout). skill/references/entity-spec.md wireframe row: drop the ''avoid tag attributes'' clause. skill/references/commands.md --wireframe flag: same. Add one short paragraph or mini table summarising the most common tags and their effect.'
id: TSK-0090
kind: task
name: Drop attribute-free UIML caveat from skill docs
objective: Skill docs guide agents to use UIML attributes freely now that the lexer fix is shipped
plan_phase: phase_5
plan_ref: uiml-wireframe-render
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: drop-attribute-free-uiml-caveat-from-skill-docs-wy5s
task_status: completed
updated_at: "2026-04-15T03:11:21Z"
---
