---
id: TSK-0043
kind: task
name: Tweak PhaseTaskList phase header rendering
slug: tweak-phasetasklist-phase-header-rendering-yzj2
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: spa-shows-plan-and-task-views-kuuc
      type: references
updated_at: "2026-04-16T01:33:38Z"
task_status: completed
priority: medium
objective: 'Each phase row in PhaseTaskList shows a layered/flag icon (not the same circle used by tasks) and the phase name is prefixed with ''Phase N: ''.'
details: 'web/src/components/PhaseTaskList.tsx: replace the leading circle glyph (○/●/✓) on the PHASE header row with a distinct phase icon — use a flag/layers Lucide icon or a styled square that visually separates phases from their nested task rows. Prefix each phase name with ''Phase N: '' where N is the 1-based ordinal among top-level phases. Status is still conveyed by color or a small status pill so the icon swap doesn''t lose information. Task rows below each phase keep their existing ○/●/✓/✗ status icons unchanged.'
acceptance: 'Opening the Tasks tab on a plan shows phases as ''Phase 1: <name>'', ''Phase 2: <name>'' etc., with a flag/layers icon in front of the phase name (not the task circle), and tasks underneath still render with the standard status icons.'
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/PhaseTaskList.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:25:31Z"
completed_at: "2026-04-15T15:10:07Z"
---
