---
id: REQ-0389
kind: requirement
name: 'Flow steps: plan lifecycle flows'
slug: flow-steps-plan-lifecycle-flows-hwas
relationships:
    - target: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      type: derives_from
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:50:36Z"
statement: The syde design model shall trace plan lifecycle flows back to the Flow-steps plan via this scoped requirement.
req_type: constraint
priority: must
verification: sync check reports no cap violation on this bucket and each flow references this requirement instead of the parent
source: plan
requirement_status: active
rationale: Cap compliance for the parent approved-plan requirement; this bucket holds 10 plan lifecycle flows.
audited_overlaps:
    - slug: flow-steps-entity-operation-flows
      distinction: plan-lifecycle-flows bucket vs entity-operation-flows bucket — both flow-kind children of Flow-steps but cover different domains (plans vs entities)
    - slug: flow-steps-dashboard-browsing-flows
      distinction: plan-lifecycle-flows bucket vs dashboard-browsing-flows bucket — both flow-kind children of Flow-steps but cover different user goals
    - slug: clear-all-sync-check-concept-redesign-tasks
      distinction: flow-steps split-child for plan-lifecycle flows vs Clear-all sync-check split-child for concept redesign tasks — different parent plans, different bucket criteria
    - slug: clear-all-sync-check-audit-and-overlap-tasks
      distinction: flow-steps split-child for plan-lifecycle flows vs Clear-all sync-check split-child for audit/overlap tasks — different parent plans, different bucket criteria
    - slug: flow-steps-flow-authoring-tasks
      distinction: plan-lifecycle-flows bucket scopes plan-CRUD flow entities; authoring-tasks bucket scopes flow-construction work — different kinds of children
    - slug: flow-steps-chart-and-doc-tasks
      distinction: plan-lifecycle-flows bucket scopes plan-CRUD flow entities; chart-and-doc-tasks bucket scopes implementation work — different children of Flow-steps
    - slug: flow-steps-entity-operation-flows-f6gc
      distinction: Scopes plan lifecycle flows, whereas target scopes entity operation flows under the same plan.
    - slug: flow-steps-dashboard-browsing-flows-nnjk
      distinction: Scopes plan lifecycle flows, whereas target scopes dashboard browsing flows under the same plan.
    - slug: clear-all-sync-check-concept-redesign-tasks-hglg
      distinction: Traces Flow-steps lifecycle flows; target traces concept-redesign tasks under the Clear-all-sync-check plan.
    - slug: clear-all-sync-check-audit-and-overlap-tasks-8lrc
      distinction: Traces Flow-steps lifecycle flows; target traces audit and overlap tasks under the Clear-all-sync-check plan.
    - slug: flow-steps-flow-authoring-tasks-tkxi
      distinction: Scopes plan lifecycle flows, whereas target scopes flow-authoring cleanup tasks under the same plan.
    - slug: flow-steps-chart-and-doc-tasks-39lu
      distinction: Scopes plan lifecycle flows, whereas target scopes chart and doc cleanup tasks under the same plan.
---
