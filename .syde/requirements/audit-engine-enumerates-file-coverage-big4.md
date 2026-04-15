---
id: REQ-0040
kind: requirement
name: Audit Engine Enumerates File Coverage
slug: audit-engine-enumerates-file-coverage-big4
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:21Z"
statement: When invoked for coverage, the audit engine shall return a mapping from every non-ignored tree file path to the set of owning entities.
req_type: functional
priority: must
verification: inspection of FileCoverage in internal/audit/orphans.go
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: File coverage drives the syde files coverage command and orphan detection.
---
