---
id: TSK-0111
kind: task
name: Show relationship total in detail panel
slug: show-relationship-total-in-detail-panel-zg66
relationships:
    - target: show-dashboard-relationship-totals-fh8a
      type: belongs_to
    - target: dashboard-must-show-relationship-totals-clearly-rc5w
      type: references
updated_at: "2026-04-15T08:19:11Z"
task_status: completed
priority: high
objective: Expose total relationship count and clearer collapsed group labels in EntityDetail.
details: Edit RelationshipsSection and RelationshipGroup in web/src/components/EntityDetail.tsx so the title includes the total and group rows read as linked entity counts.
acceptance: The Relationships section title includes the total count and collapsed groups show count plus kind label clearly.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/EntityDetail.tsx
plan_ref: show-dashboard-relationship-totals-fh8a
plan_phase: phase_1
created_at: "2026-04-15T08:18:53Z"
completed_at: "2026-04-15T08:19:11Z"
---
