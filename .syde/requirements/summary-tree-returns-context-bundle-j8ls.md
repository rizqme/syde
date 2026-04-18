---
id: REQ-0190
kind: requirement
name: Summary Tree Returns Context Bundle
slug: summary-tree-returns-context-bundle-j8ls
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:06Z"
statement: When ContextBundle is requested for a node, the summary tree shall return the ancestor breadcrumb, the node summary, and for files the inlined content capped at 64 KiB.
req_type: functional
priority: must
verification: unit test of ContextBundle in internal/tree/context.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: One-call context bundles replace ad-hoc file reads for agents.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:06Z"
---
