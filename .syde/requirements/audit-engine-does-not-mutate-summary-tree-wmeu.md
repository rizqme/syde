---
id: REQ-0054
kind: requirement
name: Audit Engine Does Not Mutate Summary Tree
slug: audit-engine-does-not-mutate-summary-tree-wmeu
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:05Z"
statement: The audit engine shall not mutate the summary tree state during an audit run.
req_type: constraint
priority: must
verification: code review of internal/audit for tree read-only access
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Tree mutation during audit would obscure stale detection.
audited_overlaps:
    - slug: audit-engine-does-not-mutate-store-ka7w
      distinction: 'Different target: this requirement forbids mutating the summary tree, while the other forbids mutating entities in the store.'
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:05Z"
---
