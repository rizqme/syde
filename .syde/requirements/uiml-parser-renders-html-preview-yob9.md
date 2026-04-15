---
id: REQ-0200
kind: requirement
name: UIML Parser Renders HTML Preview
slug: uiml-parser-renders-html-preview-yob9
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:22Z"
statement: When RenderHTML is called with a UIML AST, the UIML parser shall produce a self-contained Tailwind-styled HTML document.
req_type: functional
priority: must
verification: unit test of RenderHTML in internal/uiml/render_html.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Browser previews complement the ASCII renderer for dashboard use.
---
