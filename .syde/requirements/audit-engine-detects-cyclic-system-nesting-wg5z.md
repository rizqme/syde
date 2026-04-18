---
id: REQ-0050
kind: requirement
name: Audit Engine Detects Cyclic System Nesting
slug: audit-engine-detects-cyclic-system-nesting-wg5z
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:38:02Z"
statement: If the belongs_to graph of systems or the depends_on graph of components contains a cycle, then the audit engine shall report a cycle finding with an arrow-joined name path.
req_type: functional
priority: must
verification: inspection of internal/audit/cycles.go DFS logic
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: Cycles in hierarchy or dependency graphs break traversal and must be surfaced.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:38:02Z"
---
