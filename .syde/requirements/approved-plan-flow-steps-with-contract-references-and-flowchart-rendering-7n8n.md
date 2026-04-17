---
id: REQ-0345
kind: requirement
name: 'Approved plan: Flow steps with contract references and flowchart rendering'
slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
relationships:
    - target: flow-steps-with-contract-references-and-flowchart-rendering-m0b5
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
    - target: syde
      type: belongs_to
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T10:46:37Z"
statement: The syde flow entity shall carry an ordered step list where each step references a contract by slug, with the dashboard rendering the list as a connected flowchart and a per-user-goal decomposition replacing the catch-all flow.
req_type: functional
priority: must
verification: FlowEntity.Steps exists, the flow detail panel renders a chart, and every active contract is referenced by at least one flow step
source: plan
source_ref: plan:flow-steps-with-contract-references-and-flowchart-rendering-m0b5
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-16T10:30:53Z"
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: Flow entity schema with ordered contract-referenced steps and flowchart rendering vs. plan entity sync-check completion gate — different entity kind and different command.
    - slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      distinction: Flow entity step-list redesign plus dashboard flowchart rendering vs. CLI phase auto-complete scoping bug during syde task done — different entity kind and different subsystem.
    - slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      distinction: Flow entity step schema plus dashboard flowchart rendering vs. plan detail API bare-slug task resolution fallback — different entity kind and different API endpoint.
    - slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      distinction: Flow entity structural redesign with per-goal decomposition and flowchart rendering vs. plan authoring audit warnings on requirement-lane size and TF-IDF statement overlap — different entity kind and different concern.
    - slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      distinction: Flow entity ordered steps and dashboard rendering vs. syde add requirement TF-IDF overlap --audited acknowledgement — different entity kind, different CLI command, different concern.
---
