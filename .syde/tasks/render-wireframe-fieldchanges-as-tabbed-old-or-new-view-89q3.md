---
id: TSK-0049
kind: task
name: Render wireframe field_changes as tabbed Old or New view
slug: render-wireframe-fieldchanges-as-tabbed-old-or-new-view-89q3
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: ExtendedFieldDiff renders the wireframe field as a two-tab switcher (Current | Proposed) with each tab showing the rendered wireframe HTML, instead of the side-by-side UIML code block.
details: 'web/src/components/ExtendedFieldDiff.tsx: detect when the field key is ''wireframe'' and the current_values + proposed have wireframe_html siblings. In that case, render a tab switcher above the diff: Current tab = dangerouslySetInnerHTML current_values.wireframe_html, Proposed tab = dangerouslySetInnerHTML proposed_values_html.wireframe. All other field_changes entries continue to render in the existing side-by-side table. UIML source code stays available via a ''view source'' toggle inside each tab.'
acceptance: Opening the Plan View Screen Extended card with a wireframe field_changes shows tabs (Current | Proposed) with rendered wireframes; clicking between tabs swaps the rendered HTML.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/ExtendedFieldDiff.tsx
    - web/src/lib/api.ts
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:29:36Z"
completed_at: "2026-04-15T15:09:37Z"
---
