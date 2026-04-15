---
id: REQ-0197
kind: requirement
name: UIML Parser Builds Typed AST
slug: uiml-parser-builds-typed-ast-abcw
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:15Z"
statement: When parsing succeeds, the UIML parser shall build a typed AST of Node values whose Kind matches one of the entries in ValidTags.
req_type: functional
priority: must
verification: unit test of Parse in internal/uiml/parser.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: A typed AST constrains downstream renderers to known UIML kinds.
---
