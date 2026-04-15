---
acceptance: Dragging a concept card from its border or attribute area moves the card; clicking the name still navigates to the detail panel.
affected_entities:
    - web-spa
affected_files:
    - web/src/pages/ERD.tsx
completed_at: "2026-04-14T10:43:14Z"
created_at: "2026-04-14T10:40:01Z"
details: 'web/src/pages/ERD.tsx ConceptNode: replace the <button> wrapping the name with a <div role=''button'' onClick=... cursor-pointer>. Remove onDoubleClick from the root div. Per-attribute <Handle> elements with refs gain style={{ pointerEvents: ''none'' }}. Top/bottom node handles stay interactive. Add className=''nodrag'' to any remaining interactive element (the name div) so click-to-select works without starting a drag.'
id: TSK-0070
kind: task
name: Harden ConceptNode drag behaviour
objective: ERD node cards drag smoothly when clicking and holding any non-interactive area
plan_phase: phase_2
plan_ref: erd-polish
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: harden-conceptnode-drag-behaviour-gro6
task_status: completed
updated_at: "2026-04-14T10:43:14Z"
---
