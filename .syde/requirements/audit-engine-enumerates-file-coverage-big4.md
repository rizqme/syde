---
id: REQ-0040
kind: requirement
name: Audit Engine Enumerates File Coverage
slug: audit-engine-enumerates-file-coverage-big4
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:26Z"
statement: When invoked for coverage, the audit engine shall return a mapping from every non-ignored tree file path to the set of owning entities.
req_type: functional
priority: must
verification: inspection of FileCoverage in internal/audit/orphans.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: File coverage drives the syde files coverage command and orphan detection.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:26Z"
---
