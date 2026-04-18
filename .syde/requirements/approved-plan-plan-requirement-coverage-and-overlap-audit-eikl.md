---
id: REQ-0342
kind: requirement
name: 'Approved plan: Plan requirement coverage and overlap audit'
slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
relationships:
    - target: plan-requirement-coverage-and-overlap-audit-y33i
      type: references
      label: approved_plan
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:36Z"
statement: The syde plan authoring audit shall emit a warning when a plan's requirement lane is disproportionately small relative to its other lanes and a second warning when a new requirement's statement overlaps an existing active requirement above the TF-IDF threshold.
req_type: functional
priority: must
verification: syde plan check warns on low requirement coverage ratios and on overlap candidates that lack acknowledgement
source: plan
source_ref: plan:plan-requirement-coverage-and-overlap-audit-y33i
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-16T09:46:15Z"
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: Plan authoring audit emitting advisory warnings on requirement-lane size and TF-IDF overlap vs. syde plan complete sync-check error-gate — advisory warning vs blocking error and different command.
    - slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      distinction: Authoring-time plan audit warnings on requirement lanes vs. runtime phase auto-complete scoping fix during syde task done — authoring advisory vs runtime bug fix.
    - slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      distinction: Plan authoring audit warnings (CLI, advisory) vs. syded HTTP plan detail API bare-slug task resolution (HTTP, runtime) — different binary, different layer, different concern.
    - slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      distinction: Plan authoring audit warnings on requirement lanes vs. flow entity step-list schema and flowchart rendering — different entity kind and different concern.
    - slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      distinction: Plan-authoring audit warning on TF-IDF overlaps within a plan's requirement lane vs. syde add requirement hard-requiring --audited acknowledgement for each TF-IDF match — advisory warning vs blocking acknowledgement and different command.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:36Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:36Z"
---
