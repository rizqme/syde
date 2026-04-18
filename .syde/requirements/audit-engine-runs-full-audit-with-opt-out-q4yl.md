---
id: REQ-0034
kind: requirement
name: Audit Engine Runs Full Audit With Opt-Out
slug: audit-engine-runs-full-audit-with-opt-out-q4yl
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:36:40Z"
statement: When invoked with an options struct, the audit engine shall run the full set of audit categories while honoring per-category opt-out flags.
req_type: functional
priority: must
verification: inspection of audit.Options handling in internal/audit/audit.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Callers need to scope audits to a relevant subset for incremental checks.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:40Z"
---
