---
id: REQ-0052
kind: requirement
name: Audit Engine Does Not Mutate Store
slug: audit-engine-does-not-mutate-store-ka7w
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:45:41Z"
statement: The audit engine shall not mutate any entity in the store during an audit run.
req_type: constraint
priority: must
verification: code review of internal/audit for read-only access
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Audits must be safe to run on a live model without side effects.
audited_overlaps:
    - slug: audit-engine-does-not-mutate-summary-tree-wmeu
      distinction: 'Different target: this requirement forbids mutating store entities, while the other forbids mutating the summary tree state.'
    - slug: query-engine-does-not-mutate-store-f8c6
      distinction: 'Different component: this constrains the audit engine, while the other constrains the query engine during read operations.'
---
