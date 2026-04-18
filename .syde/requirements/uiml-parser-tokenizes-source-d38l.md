---
id: REQ-0196
kind: requirement
name: UIML Parser Tokenizes Source
slug: uiml-parser-tokenizes-source-d38l
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:36:44Z"
statement: When Parse is called with a UIML source string, the UIML parser shall tokenize the input into tag-open, tag-close, attribute, and text tokens with line tracking.
req_type: functional
priority: must
verification: unit test of lexer in internal/uiml/lexer.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: A dedicated lexer keeps the parser simple and line-accurate.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:36:44Z"
---
