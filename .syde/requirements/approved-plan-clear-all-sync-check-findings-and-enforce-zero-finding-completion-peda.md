---
id: REQ-0366
kind: requirement
name: 'Approved plan: Clear all sync check findings and enforce zero-finding completion'
slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
relationships:
    - target: clear-all-sync-check-findings-and-enforce-zero-finding-completion-nk7a
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T10:45:43Z"
statement: When syde plan complete is invoked, the syde CLI shall refuse to mark the plan completed if syde sync check reports any errors unless the author passes --force.
req_type: functional
priority: must
verification: syde plan complete blocks on non-zero sync check and skill docs document the gate
source: plan
source_ref: plan:clear-all-sync-check-findings-and-enforce-zero-finding-completion-nk7a
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T01:36:05Z"
audited_overlaps:
    - slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      distinction: Gates syde plan complete on sync check errors, versus scoping task lookup within phase auto-complete when syde task done fires — different command, different stage.
    - slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      distinction: CLI plan completion gate on sync errors vs. syded HTTP plan detail API's bare-slug fallback for task resolution — different binary and different endpoint.
    - slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      distinction: Plan-complete sync-error guard on plan entities vs. flow entity schema carrying ordered contract-referenced steps and flowchart rendering — different entity kind entirely.
    - slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      distinction: Blocks syde plan complete on sync errors vs. emits authoring-time warnings about requirement-lane size and TF-IDF overlap during plan drafting — error-gate vs advisory warning.
    - slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      distinction: Triggers on syde plan complete blocking on sync errors vs. triggers on syde add requirement requiring --audited for TF-IDF matches — different command and different subject entity.
    - slug: plan-complete-shall-require-clean-sync-check-9jcs
      distinction: approved-plan requirement records plan intent at approval time whereas the clean-sync-check rule is the behavioural gate on syde plan complete — both exist by design as intent vs implementation
---
