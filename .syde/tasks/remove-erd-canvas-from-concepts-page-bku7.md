---
id: TSK-0082
kind: task
name: Remove ERD canvas from concepts page
slug: remove-erd-canvas-from-concepts-page-bku7
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: Concepts page has no ERD toggle or canvas; standard 2-column inbox only
details: Remove ERD rendering code, List/ERD toggle, and any ERD-specific components
acceptance: Concepts page shows standard list + detail layout
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/App.tsx
    - web/src/components/EntityDetail.tsx
    - web/package.json
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_3
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T08:16:25Z"
---
