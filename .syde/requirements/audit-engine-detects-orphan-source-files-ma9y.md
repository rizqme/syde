---
id: REQ-0036
kind: requirement
name: Audit Engine Detects Orphan Source Files
slug: audit-engine-detects-orphan-source-files-ma9y
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:16Z"
statement: The audit engine shall detect non-ignored tree files that have no owning entity and emit them as orphan findings.
req_type: functional
priority: must
verification: inspection of internal/audit/orphans.go Orphans() and orphanFindings
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Orphan source files indicate gaps between the codebase and the design model.
---
