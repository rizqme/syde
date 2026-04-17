---
id: REQ-0365
kind: requirement
name: 'Approved plan: Fix task resolution in plan detail API'
slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
relationships:
    - target: fix-task-resolution-in-plan-detail-api-otih
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T10:46:19Z"
statement: When the syded plan detail API cannot resolve a task by exact slug match, the API shall fall back to bare-slug comparison so plans that reference tasks by bare name still resolve.
req_type: functional
priority: must
verification: a plan detail fetch for a plan referencing tasks by bare slug returns the expected task payload
source: plan
source_ref: plan:fix-task-resolution-in-plan-detail-api-otih
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-16T11:49:56Z"
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: syded HTTP plan detail API task-resolution fallback vs. CLI syde plan complete gate on sync check errors — different binary (syded vs syde) and different stage.
    - slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      distinction: HTTP plan detail API bare-slug fallback for task resolution (read path) vs. CLI phase auto-complete scoping during syde task done (write path) — different subsystems.
    - slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      distinction: HTTP plan detail API task-slug resolution bug fix vs. flow entity schema redesign with contract-referenced ordered steps and flowchart rendering — different entity kind and different layer.
    - slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      distinction: HTTP API runtime task-slug resolution fallback vs. CLI authoring-time audit emitting warnings on requirement-lane size and statement overlap — runtime resolver vs authoring advisory.
    - slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      distinction: syded plan detail API bare-slug fallback for tasks vs. syde add requirement TF-IDF overlap acknowledgement — different binary, different command, different subject entity.
---
