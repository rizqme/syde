---
id: COM-0022
kind: component
name: Plan Detail Panel
slug: plan-detail-panel-nqq1
description: Inline plan detail component with markdown rendering, Plan/Tasks tabs
purpose: Match the canonical 2-column inbox UX for plans.
files:
    - web/src/components/PlanDetailPanel.tsx
relationships:
    - target: syded-dashboard-e82c
      type: belongs_to
    - target: web-spa-jy9z
      type: depends_on
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: implements
updated_at: "2026-04-17T01:48:57Z"
responsibility: Render a selected plan's design, structured changes, phases, and nested tasks inside the right column of the Plans inbox.
capabilities:
    - Render plan header with name, status, progress, and approved timestamp
    - Switch between Plan and Tasks tabs via the tab query parameter
    - Embed PlanChangesView for the structured diff
    - Embed PhaseTaskList for the Tasks tab
boundaries: Does not own routing; App.tsx does. Does not fetch the plan list; EntityList does.
---
