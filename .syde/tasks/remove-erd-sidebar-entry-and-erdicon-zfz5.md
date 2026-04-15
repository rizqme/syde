---
acceptance: Sidebar has only File Tree / Graph above the kind groups. bun run build succeeds.
affected_entities:
    - web-spa
affected_files:
    - web/src/components/Sidebar.tsx
    - web/src/components/icons.tsx
    - web/src/App.tsx
completed_at: "2026-04-14T10:24:29Z"
created_at: "2026-04-14T10:22:14Z"
details: 'web/src/components/Sidebar.tsx: drop the ERD button, drop ErdIcon from the import list. web/src/components/icons.tsx: delete the ErdIcon export. web/src/App.tsx: remove __erd__ from SPECIAL_VIEWS array and remove the __erd__ conditional render block. Keep ERD.tsx file intact — phase 3 reuses it.'
id: TSK-0064
kind: task
name: Remove __erd__ sidebar entry and ErdIcon
objective: Sidebar no longer has a standalone ERD button and __erd__ routing is gone
plan_phase: phase_2
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: remove-erd-sidebar-entry-and-erdicon-zfz5
task_status: completed
updated_at: "2026-04-14T10:24:29Z"
---
