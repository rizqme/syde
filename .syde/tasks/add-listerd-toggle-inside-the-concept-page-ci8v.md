---
acceptance: Navigating to Concepts shows List view by default with a toggle; clicking ERD renders the React Flow canvas; clicking List returns to the 2-column inbox.
affected_entities:
    - web-spa
affected_files:
    - web/src/App.tsx
completed_at: "2026-04-14T10:25:47Z"
created_at: "2026-04-14T10:22:14Z"
details: 'web/src/App.tsx: new conceptView state (''list'' | ''erd''), defaults to ''list'', resets to ''list'' when activeKind changes. When activeKind === ''concept'', render a small top-bar ConceptViewTabs component above the content area. When conceptView === ''list'' render the existing 2-column; when ''erd'' render <ERD /> filling the main area. Clicking a node in ERD navigates to the entity detail AND sets conceptView=''list'' so the detail is visible. Toggle is not persisted across reloads.'
id: TSK-0065
kind: task
name: Add List/ERD toggle inside the Concept page
objective: Clicking 'Concepts' shows the list+detail inbox by default; a segmented toggle switches the main area into ERD mode
plan_phase: phase_3
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-listerd-toggle-inside-the-concept-page-ci8v
task_status: completed
updated_at: "2026-04-14T10:25:47Z"
---
