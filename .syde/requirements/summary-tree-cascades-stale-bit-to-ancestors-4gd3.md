---
id: REQ-0187
kind: requirement
name: Summary Tree Cascades Stale Bit To Ancestors
slug: summary-tree-cascades-stale-bit-to-ancestors-4gd3
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:44Z"
statement: When a file's hash changes during a scan, the summary tree shall mark the file stale and cascade the stale bit up to every ancestor directory.
req_type: functional
priority: must
verification: unit test of Scan in internal/tree/scan.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Cascading staleness lets agents walk leaves first while still knowing parents need attention.
---
