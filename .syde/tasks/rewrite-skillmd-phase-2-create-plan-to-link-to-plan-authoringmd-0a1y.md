---
id: TSK-0266
kind: task
name: Rewrite SKILL.md Phase 2 (CREATE PLAN) to link to plan-authoring.md
slug: rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd-0a1y
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: plan-tasks-shall-decompose-work-into-bite-sized-checkbox-steps-jv3l
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: medium
objective: 'SKILL.md Phase 2 enforces the new discipline by explicit reference: Files section required, bite-sized checkbox steps required, No Placeholders blacklist enforced, Self-Review checklist mandatory, Execution Handoff prompt required.'
details: 'Locate the existing ''### Phase 2: CREATE PLAN'' section in skill/SKILL.md. Prepend a short paragraph pointing the author at references/plan-authoring.md for the full discipline. Add five bullet-level enforcement items summarising the five rules. Keep the existing step-by-step plan-create / add-phase / add-change / task-create guidance intact (it is syde-specific) but make clear that the NEW discipline is overlaid on top.'
acceptance: skill/SKILL.md Phase 2 section links to references/plan-authoring.md; enumerates the five new discipline rules; does not duplicate the full doc content (links only).
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:24:42Z"
completed_at: "2026-04-18T09:35:05Z"
---
