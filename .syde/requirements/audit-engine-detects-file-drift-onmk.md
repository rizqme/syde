---
id: REQ-0038
kind: requirement
name: Audit Engine Detects File Drift
slug: audit-engine-detects-file-drift-onmk
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:29Z"
statement: When an owned file's mtime is newer than the owning entity's UpdatedAt timestamp, the audit engine shall emit a drift warning finding.
req_type: functional
priority: must
verification: inspection of driftFindings in internal/audit/orphans.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Drift signals that the design model needs resync with the code.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:29Z"
---
