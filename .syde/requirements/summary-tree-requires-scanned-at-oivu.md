---
id: REQ-0113
kind: requirement
name: Summary Tree Requires Scanned At
slug: summary-tree-requires-scanned-at-oivu
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:28Z"
statement: The syde CLI shall require a scanned_at timestamp on every persisted summary tree instance.
req_type: constraint
priority: must
verification: unit test loading a tree.yaml without scanned_at
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: scanned_at anchors diff detection and status reporting.
---
