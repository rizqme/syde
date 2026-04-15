---
id: REQ-0161
kind: requirement
name: Scan Helpers Map Files To Components
slug: scan-helpers-map-files-to-components-6avp
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:23Z"
statement: When sync coverage is invoked, the scan helpers shall compare scanned directories against component file globs and report mapped versus unmapped directories.
req_type: functional
priority: must
verification: inspection of coverage report generation in internal/scan/coverage.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: File-to-component coverage is the legacy completeness signal.
---
