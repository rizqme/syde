---
acceptance: A fresh install-skill lands the updated files; agents reading the skill can copy the example and it parses through uiml.Parse.
affected_entities:
    - skill-installer
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/references/commands.md
completed_at: "2026-04-14T11:34:03Z"
created_at: "2026-04-14T11:23:14Z"
details: 'skill/SKILL.md contract rules: add ''screen'' to contract_kind enum list with ''render'' as the paired interaction_pattern; new subsection with a Concepts List example showing UIML with sidebar+main+detail. skill/references/entity-spec.md Contract table: add ''wireframe'' row with syntax pointer. skill/references/commands.md: add --wireframe to contract flag reference.'
id: TSK-0078
kind: task
name: Document screen + render + --wireframe in the skill
objective: SKILL.md, entity-spec.md, commands.md describe screen contracts with a worked UIML example
plan_phase: phase_3
plan_ref: screen-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: document-screen-render-wireframe-in-the-skill-quur
task_status: completed
updated_at: "2026-04-14T11:34:03Z"
---
