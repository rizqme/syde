---
id: TSK-0035
kind: task
name: Render wireframe HTML in PlanChangesView Extended and New cards
slug: render-wireframe-html-in-planchangesview-extended-and-new-cards-lxce
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: PlanChangesView shows the rendered wireframe HTML inline for any Extended card targeting a screen contract and any New card whose draft.contract_kind === 'screen'.
details: ExtendedFieldDiff (or PlanChangesView itself) checks for current_values.wireframe_html and renders it via <div dangerouslySetInnerHTML={...} /> at the top of the card body. NewContractDraftView's screen branch switches from the placeholder <pre> block to dangerouslySetInnerHTML when draft.wireframe_html is present, falling back to the <pre> source when only draft.wireframe is set. Reuse the styling used by the existing EntityDetail screen wireframe (look at how EntityDetail renders wireframe_html for screen contracts and copy the wrapper).
acceptance: Opening the Plans Inbox 2-Column Layout plan in the dashboard renders the rendered Plan View Screen wireframe inside the Contracts Extended card.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/PlanChangesView.tsx
    - web/src/components/NewContractDraftView.tsx
    - web/src/lib/api.ts
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:07:14Z"
completed_at: "2026-04-15T15:09:37Z"
---
