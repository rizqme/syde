---
id: REQ-0341
kind: requirement
name: 'Approved plan: Fix phase auto-completion cross-plan collision'
slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
relationships:
    - target: fix-phase-auto-completion-cross-plan-collision-h79u
      type: references
      label: approved_plan
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:08Z"
statement: When syde task done completes the last task in a phase, the phase auto-complete logic shall scope task lookup to the same plan and match tasks by entity slug rather than by re-slugifying the task name.
req_type: functional
priority: must
verification: a crafted two-plan fixture with same-named tasks auto-completes only the owning phase
source: plan
source_ref: plan:fix-phase-auto-completion-cross-plan-collision-h79u
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-16T05:15:50Z"
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: Scopes phase auto-complete task lookup to the same plan on syde task done vs. gates syde plan complete on sync check errors — different command and different stage.
    - slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      distinction: CLI-side phase auto-complete scoping during syde task done vs. syded HTTP plan detail API falling back to bare-slug for task resolution — CLI write path vs HTTP read path.
    - slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      distinction: Phase auto-complete bug fix in syde task done on plan entities vs. flow entity structural redesign with ordered contract-referenced steps — different entity kind and different concern.
    - slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      distinction: Runtime phase auto-complete logic during syde task done vs. authoring-time advisory warnings about requirement-lane size and statement overlap during plan drafting — runtime bug fix vs audit warning.
    - slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      distinction: Fires on syde task done to correctly scope phase completion vs. fires on syde add requirement requiring --audited for TF-IDF overlaps — different command and different subject entity kind.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:08Z"
---
