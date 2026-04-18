---
id: REQ-0201
kind: requirement
name: UIML Parser Validates Structural Rules
slug: uiml-parser-validates-structural-rules-pjt0
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:36:55Z"
statement: When Validate is called with a UIML source, the UIML parser shall report semantic errors including unknown tags, missing columns in tables, and misuse of variant or direction attributes.
req_type: functional
priority: must
verification: unit test of Validate in internal/uiml/validate.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Structural validation catches authoring mistakes before persistence.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:36:55Z"
---
