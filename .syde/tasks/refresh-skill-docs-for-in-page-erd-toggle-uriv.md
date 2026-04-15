---
acceptance: grep 'sidebar.*ERD' skill/ returns no hits. Web SPA component description mentions the in-page ERD toggle.
affected_entities:
    - skill-installer
    - web-spa
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
completed_at: "2026-04-14T10:30:17Z"
created_at: "2026-04-14T10:22:14Z"
details: 'skill/SKILL.md Concept rules: replace ''open the dashboard and click ERD in the sidebar'' with ''open the dashboard, click Concepts, toggle to ERD view''. skill/references/entity-spec.md mirror. Also update the Web SPA component summary (syde update web-spa) to describe the in-page toggle and the ERD node card semantics (name + description + attribute refs).'
id: TSK-0066
kind: task
name: Refresh skill docs for in-page ERD toggle
objective: SKILL.md and entity-spec.md describe the concept view-mode toggle instead of a sidebar item
plan_phase: phase_4
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: refresh-skill-docs-for-in-page-erd-toggle-uriv
task_status: completed
updated_at: "2026-04-14T10:30:17Z"
---
