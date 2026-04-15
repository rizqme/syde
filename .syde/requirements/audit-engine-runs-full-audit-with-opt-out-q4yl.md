---
id: REQ-0034
kind: requirement
name: Audit Engine Runs Full Audit With Opt-Out
slug: audit-engine-runs-full-audit-with-opt-out-q4yl
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:13Z"
statement: When invoked with an options struct, the audit engine shall run the full set of audit categories while honoring per-category opt-out flags.
req_type: functional
priority: must
verification: inspection of audit.Options handling in internal/audit/audit.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Callers need to scope audits to a relevant subset for incremental checks.
---
