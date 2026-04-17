---
id: TSK-0015
kind: task
name: ExtendedFieldDiff side-by-side view
slug: extendedfielddiff-side-by-side-view-c01p
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: web/src/components/ExtendedFieldDiff.tsx renders a side-by-side old-vs-new view for every declared field in an ExtendedChange.
details: 'Takes Current map[string]any and FieldChanges map[string]any. For each key in FieldChanges: render field name, current value (monospace block), proposed value (monospace block). Multi-line strings are unified-diff-highlighted. Missing keys in Current render as italic ''(not set)''. A DELETE sentinel in FieldChanges renders the proposed side as red strikethrough ''(remove)''.'
acceptance: An Extended change with field_changes containing responsibility + boundaries renders both fields side-by-side.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:19:14Z"
---
