---
id: REQ-0200
kind: requirement
name: UIML Parser Renders HTML Preview
slug: uiml-parser-renders-html-preview-yob9
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:37:10Z"
statement: When RenderHTML is called with a UIML AST, the UIML parser shall produce a self-contained Tailwind-styled HTML document.
req_type: functional
priority: must
verification: unit test of RenderHTML in internal/uiml/render_html.go
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Browser previews complement the ASCII renderer for dashboard use.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:37:10Z"
---
