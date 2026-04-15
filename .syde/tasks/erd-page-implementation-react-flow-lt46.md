---
acceptance: Visiting the ERD sidebar item renders all concepts as draggable cards with attributes visible and relates_to edges labelled.
affected_entities:
    - web-spa
completed_at: "2026-04-14T09:58:19Z"
created_at: "2026-04-14T09:44:44Z"
details: 'New file web/src/pages/ERD.tsx (add to web-spa --file list at completion). Fetch concept list via useApi/useEntities. Build nodes array: id=entity slug, data={name, attributes, actions}, position in a simple grid (index % 5 * 320, index / 5 * 220). Build edges array by iterating concept.relationships where type==relates_to; id=source-target, source=source slug, target=target slug, label=rel.label. Wrap in ReactFlowProvider. Custom ConceptNode with <Handle> top and bottom. Import ''@xyflow/react/dist/style.css''.'
id: TSK-0062
kind: task
name: ERD page implementation (React Flow)
objective: web/src/pages/ERD.tsx exists and renders the ERD correctly against live data from /api/<proj>/list?kind=concept
plan_phase: phase_4
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: erd-page-implementation-react-flow-lt46
task_status: completed
updated_at: "2026-04-14T09:58:19Z"
---
