---
id: TSK-0024
kind: task
name: Plans Inbox column page
slug: plans-inbox-column-page-gg8w
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: Plans get an inbox-style column page at /__plan__ matching the Systems/Components/Contracts/Concepts/Flows inbox layout.
details: 'Add web/src/pages/PlansInboxScreen.tsx mirroring the existing *InboxScreen pages: left column lists every plan (name + status badge + progress bar + short objective), selecting one navigates to /__plan__/<slug>. Filter bar supports status (draft/approved/completed). Uses the existing EntityFilterBar + KindBadge components. Wire into Sidebar as a primary nav item that replaces the removed Tasks item.'
acceptance: Opening /__plan__ shows every plan as a selectable row with status and progress; selecting a plan navigates to the detail page.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:44:58Z"
completed_at: "2026-04-15T12:14:54Z"
---
