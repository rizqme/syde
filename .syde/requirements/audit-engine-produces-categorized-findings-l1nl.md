---
id: REQ-0033
kind: requirement
name: Audit Engine Produces Categorized Findings
slug: audit-engine-produces-categorized-findings-l1nl
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:07Z"
statement: The audit engine shall produce categorized findings covering entity field validation, relationship integrity, cycles, tree file references, orphan detection, and file drift.
req_type: functional
priority: must
verification: inspection of internal/audit/audit.go Run() aggregation
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Centralizing health-check outputs avoids re-implementing the same checks across validate, sync check, and files commands.
---
