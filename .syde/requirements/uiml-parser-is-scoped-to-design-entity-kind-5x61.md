---
id: REQ-0202
kind: requirement
name: UIML Parser Is Scoped To Design Entity Kind
slug: uiml-parser-is-scoped-to-design-entity-kind-5x61
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:27Z"
statement: The UIML parser shall be used only for entities of the design kind and shall not parse other entity bodies.
req_type: constraint
priority: must
verification: code review of internal/uiml callers
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: UIML is the body language for designs only.
---
