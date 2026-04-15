---
acceptance: 'Concept page: no dedicated tab row; toggle visible at top-right of the main area in both list and ERD mode; X close button absent from concept detail panel.'
affected_entities:
    - web-spa
affected_files:
    - web/src/App.tsx
    - web/src/components/EntityDetail.tsx
completed_at: "2026-04-14T10:42:21Z"
created_at: "2026-04-14T10:40:01Z"
details: 'web/src/App.tsx: remove the border-b dedicated row. Wrap the main concept area in ''relative''. Add an absolute-positioned div (top-4 right-4 z-10) containing the segmented toggle. In list mode the toggle floats above the detail pane; in ERD mode above the canvas. EntityDetail.tsx: add optional hideClose prop on DetailShell; when activeKind === ''concept'', pass hideClose={true} from App.tsx so the X button does not render and the toggle occupies its spot.'
id: TSK-0069
kind: task
name: Move toggle to top-right overlay, hide X in concept detail
objective: Concept page renders the List/ERD toggle at the top-right of the main area; EntityDetail X button is hidden on kind=concept
plan_phase: phase_1
plan_ref: erd-polish
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: move-toggle-to-top-right-overlay-hide-x-in-concept-detail-mk99
task_status: completed
updated_at: "2026-04-14T10:42:21Z"
---
