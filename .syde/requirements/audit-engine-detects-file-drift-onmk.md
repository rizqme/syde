---
id: REQ-0038
kind: requirement
name: Audit Engine Detects File Drift
slug: audit-engine-detects-file-drift-onmk
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:18Z"
statement: When an owned file's mtime is newer than the owning entity's UpdatedAt timestamp, the audit engine shall emit a drift warning finding.
req_type: functional
priority: must
verification: inspection of driftFindings in internal/audit/orphans.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Drift signals that the design model needs resync with the code.
---
