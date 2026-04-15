---
id: REQ-0201
kind: requirement
name: UIML Parser Validates Structural Rules
slug: uiml-parser-validates-structural-rules-pjt0
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:25Z"
statement: When Validate is called with a UIML source, the UIML parser shall report semantic errors including unknown tags, missing columns in tables, and misuse of variant or direction attributes.
req_type: functional
priority: must
verification: unit test of Validate in internal/uiml/validate.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Structural validation catches authoring mistakes before persistence.
---
