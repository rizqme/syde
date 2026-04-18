---
id: REQ-0199
kind: requirement
name: UIML Parser Renders ASCII Preview
slug: uiml-parser-renders-ascii-preview-5pme
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:37:28Z"
statement: When RenderASCII is called with a UIML AST and width, the UIML parser shall emit a box-drawing terminal preview of the design.
req_type: functional
priority: must
verification: unit test of RenderASCII in internal/uiml/render_ascii.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: ASCII previews let agents inspect designs in the terminal.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:37:28Z"
---
