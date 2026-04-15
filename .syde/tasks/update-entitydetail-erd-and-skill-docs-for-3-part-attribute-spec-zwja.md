---
acceptance: Dashboard renders attributes as two columns (name + description); syde sync check clean; ERD shows back-filled concepts with descriptions (no type).
affected_entities:
    - web-spa
    - skill-installer
affected_files:
    - web/src/lib/api.ts
    - web/src/components/EntityDetail.tsx
    - web/src/pages/ERD.tsx
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/references/commands.md
completed_at: "2026-04-14T10:51:43Z"
created_at: "2026-04-14T10:40:01Z"
details: 'web/src/lib/api.ts: drop type from the attribute TS type. web/src/components/EntityDetail.tsx: concept ParamTable no longer passes type. web/src/pages/ERD.tsx: already doesn''t render type, just verify the ConceptAttribute TS interface. skill/SKILL.md: Concept rules section rewrites every --attribute example to ''name|description[|refs]''. skill/references/entity-spec.md: attribute spec table updated. skill/references/commands.md: flag help updated. Re-backfill Entity/Plan/Task/Plan Phase/Relationship/Decision/Learning/FileRef/Skill/Summary Tree/Tree Node concepts with the new 3-part specs.'
id: TSK-0072
kind: task
name: Update EntityDetail, ERD, and skill docs for 3-part attribute spec
objective: Dashboard and docs reflect the typeless 3-part spec; back-filled concepts rewritten without type
plan_phase: phase_3
plan_ref: erd-polish
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: update-entitydetail-erd-and-skill-docs-for-3-part-attribute-spec-zwja
task_status: completed
updated_at: "2026-04-14T10:51:43Z"
---
