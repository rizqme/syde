---
id: REQ-0056
kind: requirement
name: Audit Engine Does Not Format Output
slug: audit-engine-does-not-format-output-g8hh
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:36:56Z"
statement: The audit engine shall not format findings for display and shall leave rendering to its callers.
req_type: constraint
priority: must
verification: inspection of audit package API surface
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Separating data from presentation lets CLI and HTTP API consume the same findings.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:56Z"
---
