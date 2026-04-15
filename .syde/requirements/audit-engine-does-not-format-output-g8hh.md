---
id: REQ-0056
kind: requirement
name: Audit Engine Does Not Format Output
slug: audit-engine-does-not-format-output-g8hh
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:33Z"
statement: The audit engine shall not format findings for display and shall leave rendering to its callers.
req_type: constraint
priority: must
verification: inspection of audit package API surface
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Separating data from presentation lets CLI and HTTP API consume the same findings.
---
