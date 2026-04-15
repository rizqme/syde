---
id: REQ-0199
kind: requirement
name: UIML Parser Renders ASCII Preview
slug: uiml-parser-renders-ascii-preview-5pme
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:20Z"
statement: When RenderASCII is called with a UIML AST and width, the UIML parser shall emit a box-drawing terminal preview of the design.
req_type: functional
priority: must
verification: unit test of RenderASCII in internal/uiml/render_ascii.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: ASCII previews let agents inspect designs in the terminal.
---
