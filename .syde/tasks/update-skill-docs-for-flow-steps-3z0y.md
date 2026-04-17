---
id: TSK-0073
kind: task
name: Update skill docs for flow steps
slug: update-skill-docs-for-flow-steps-3z0y
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: SKILL.md documents the --step flag and flow step authoring
details: Update the flow section in SKILL.md to document structured steps, the --step flag format, and the audit rules
acceptance: grep 'step' skill/SKILL.md returns flow step documentation
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_5
created_at: "2026-04-16T09:23:43Z"
completed_at: "2026-04-16T10:57:39Z"
---
