---
id: REQ-0160
kind: requirement
name: Scan Helpers Produce Scan Guide
slug: scan-helpers-produce-scan-guide-81e6
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:38:00Z"
statement: When syde sync is invoked, the scan helpers shall produce a ScanGuide containing file count, language histogram, significant source directories, and key files.
req_type: functional
priority: must
verification: inspection of ScanGuide emission in internal/scan/guide.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Legacy scan commands rely on this guide for coverage reporting.
---
