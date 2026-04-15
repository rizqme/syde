---
id: REQ-0193
kind: requirement
name: Summary Tree Persists To tree yaml Atomically
slug: summary-tree-persists-to-tree-yaml-atomically-jg5a
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:02Z"
statement: When Save is called, the summary tree shall write .syde/tree.yaml using an atomic rename so partial writes cannot corrupt the file.
req_type: functional
priority: must
verification: unit test of Save in internal/tree/store.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Atomic rename protects against interrupted scans.
---
