---
id: REQ-0202
kind: requirement
name: UIML Parser Is Scoped To Design Entity Kind
slug: uiml-parser-is-scoped-to-design-entity-kind-5x61
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:36:41Z"
statement: The UIML parser shall be used only for entities of the design kind and shall not parse other entity bodies.
req_type: constraint
priority: must
verification: code review of internal/uiml callers
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: UIML is the body language for designs only.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:36:41Z"
---
