---
id: REQ-0050
kind: requirement
name: Audit Engine Detects Cyclic System Nesting
slug: audit-engine-detects-cyclic-system-nesting-wg5z
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:23Z"
statement: If the belongs_to graph of systems or the depends_on graph of components contains a cycle, then the audit engine shall report a cycle finding with an arrow-joined name path.
req_type: functional
priority: must
verification: inspection of internal/audit/cycles.go DFS logic
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Cycles in hierarchy or dependency graphs break traversal and must be surfaced.
---
