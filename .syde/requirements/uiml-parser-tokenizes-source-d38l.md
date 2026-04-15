---
id: REQ-0196
kind: requirement
name: UIML Parser Tokenizes Source
slug: uiml-parser-tokenizes-source-d38l
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:12Z"
statement: When Parse is called with a UIML source string, the UIML parser shall tokenize the input into tag-open, tag-close, attribute, and text tokens with line tracking.
req_type: functional
priority: must
verification: unit test of lexer in internal/uiml/lexer.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: A dedicated lexer keeps the parser simple and line-accurate.
---
