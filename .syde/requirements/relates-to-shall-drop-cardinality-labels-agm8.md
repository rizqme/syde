---
id: REQ-0376
kind: requirement
name: Relates-to shall drop cardinality labels
slug: relates-to-shall-drop-cardinality-labels-agm8
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T08:25:33Z"
statement: The syde audit engine shall not validate cardinality labels on relates_to relationships between concept entities.
req_type: constraint
priority: must
verification: relates_to without cardinality succeeds
source: plan
requirement_status: active
rationale: Prose labels more natural
---
