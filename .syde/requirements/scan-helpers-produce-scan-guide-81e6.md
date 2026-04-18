---
id: REQ-0160
kind: requirement
name: Scan Helpers Produce Scan Guide
slug: scan-helpers-produce-scan-guide-81e6
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
updated_at: "2026-04-18T09:36:56Z"
statement: When syde sync is invoked, the scan helpers shall produce a ScanGuide containing file count, language histogram, significant source directories, and key files.
req_type: functional
priority: must
verification: inspection of ScanGuide emission in internal/scan/guide.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Legacy scan commands rely on this guide for coverage reporting.
verified_against:
    scan-helpers-legacy-sa6d:
        hash: c9b19fd18480f13b89908a38caa95ff8ea66c15c73dc6cf6e551e414e71ce3ba
        at: "2026-04-18T09:36:56Z"
---
