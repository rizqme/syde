---
acceptance: bun run build succeeds. Clicking any concept in the list view shows attributes + actions. Clicking 'ERD' in the sidebar shows a React Flow canvas with every concept as a draggable card and relates_to edges labelled with cardinality.
affected_entities:
    - web-spa
affected_files:
    - web/package.json
    - web/src/components/EntityDetail.tsx
    - web/src/components/Sidebar.tsx
    - web/src/App.tsx
completed_at: "2026-04-14T09:54:48Z"
created_at: "2026-04-14T09:44:26Z"
details: 'web/: bun add @xyflow/react. EntityDetail.tsx concept case: after existing fields, render Attributes as a compact 3-column table (Name, Type, Description) and Actions as a 2-column table (Name, Description). New web/src/pages/ERD.tsx: React Flow with nodes = concepts (custom node component showing name + top-3 attributes + ''N more'' footer if truncated), edges = relates_to rendered with cardinality label as edge label. Nodes drag but don''t persist. Sidebar.tsx adds ''ERD'' button under Architecture group, routed via ''__erd__'' in App.tsx. Node styling should match existing entity card look (border, background).'
id: TSK-0060
kind: task
name: Dashboard EntityDetail panel + React Flow ERD page
objective: Concept detail shows attributes/actions; new ERD page renders every concept + relates_to visually
plan_phase: phase_4
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dashboard-entitydetail-panel-react-flow-erd-page-el0z
task_status: completed
updated_at: "2026-04-14T09:54:48Z"
---
