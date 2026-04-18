---
id: REQ-0197
kind: requirement
name: UIML Parser Builds Typed AST
slug: uiml-parser-builds-typed-ast-abcw
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:37:52Z"
statement: When parsing succeeds, the UIML parser shall build a typed AST of Node values whose Kind matches one of the entries in ValidTags.
req_type: functional
priority: must
verification: unit test of Parse in internal/uiml/parser.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: A typed AST constrains downstream renderers to known UIML kinds.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:37:52Z"
---
