---
id: REQ-0391
kind: requirement
name: 'Flow steps: dashboard browsing flows'
slug: flow-steps-dashboard-browsing-flows-nnjk
relationships:
    - target: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      type: derives_from
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:45:48Z"
statement: The syde design model shall trace dashboard browsing flows back to the Flow-steps plan via this scoped requirement.
req_type: constraint
priority: must
verification: sync check reports no cap violation on this bucket and each flow references this requirement instead of the parent
source: plan
requirement_status: active
rationale: Cap compliance for the parent approved-plan requirement; this bucket holds 10 dashboard browsing flows.
audited_overlaps:
    - slug: flow-steps-plan-lifecycle-flows
      distinction: Scopes dashboard browsing flow coverage, not plan lifecycle flow coverage under the Flow-steps plan.
    - slug: flow-steps-entity-operation-flows
      distinction: Scopes dashboard browsing flow coverage, not entity operation flow coverage under the Flow-steps plan.
    - slug: clear-all-sync-check-concept-redesign-tasks
      distinction: Traces dashboard browsing flows to the Flow-steps plan, not concept-redesign tasks under Clear-all-sync-check.
    - slug: clear-all-sync-check-audit-and-overlap-tasks
      distinction: Traces dashboard browsing flows to the Flow-steps plan, not audit and overlap tasks under Clear-all-sync-check.
    - slug: flow-steps-flow-authoring-tasks
      distinction: Scopes dashboard browsing flow coverage, not flow authoring cleanup tasks under the Flow-steps plan.
    - slug: flow-steps-chart-and-doc-tasks
      distinction: Scopes dashboard browsing flow coverage, not chart and doc cleanup tasks under the Flow-steps plan.
---
