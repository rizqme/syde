---
id: PLN-0020
kind: plan
name: Show Dashboard Relationship Totals
slug: show-dashboard-relationship-totals-fh8a
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: dashboard-must-show-relationship-totals-clearly-rc5w
      type: references
    - target: approved-plan-show-dashboard-relationship-totals-d8cm
      type: references
      label: requirement
updated_at: "2026-04-15T08:19:13Z"
plan_status: completed
background: The requirement detail API returns hundreds of inbound relationships, but the dashboard detail panel hides the total behind collapsed relationship groups.
objective: Make relationship totals obvious in the entity detail panel so large requirement link sets are discoverable.
scope: Update the Web SPA relationship section presentation only; do not change relationship API shape or graph semantics.
source: manual
created_at: "2026-04-15T08:18:43Z"
approved_at: "2026-04-15T08:18:53Z"
completed_at: "2026-04-15T08:19:13Z"
phases:
    - id: phase_1
      name: Relationship detail visibility
      status: completed
      description: Make relationship totals and grouped counts obvious in the entity detail panel.
      objective: Users can immediately see that the baseline requirement has hundreds of linked entities.
      changes: Update EntityDetail relationship section title and collapsed group row copy.
      details: Show total relationship count in the section heading and keep each collapsed group row count visually explicit.
      tasks:
        - show-relationship-total-in-detail-panel
---
