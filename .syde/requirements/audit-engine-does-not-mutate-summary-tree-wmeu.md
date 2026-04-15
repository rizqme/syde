---
id: REQ-0054
kind: requirement
name: Audit Engine Does Not Mutate Summary Tree
slug: audit-engine-does-not-mutate-summary-tree-wmeu
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:30Z"
statement: The audit engine shall not mutate the summary tree state during an audit run.
req_type: constraint
priority: must
verification: code review of internal/audit for tree read-only access
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Tree mutation during audit would obscure stale detection.
---
