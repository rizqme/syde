---
id: REQ-0032
kind: requirement
name: Component Requires Capabilities
slug: component-requires-capabilities-h52l
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:06Z"
statement: The syde CLI shall require a purpose, a responsibility, and at least one capability on every component entity instance.
req_type: constraint
priority: must
verification: integration test running syde add component without capabilities
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Components without stated capabilities cannot drive requirement derivation.
---
