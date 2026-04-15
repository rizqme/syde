---
boundaries: Scoped to the 'design' entity kind only. No runtime UI rendering.
capabilities:
    - Tokenization and AST construction
    - ASCII preview rendering (terminal)
    - HTML preview rendering (browser)
    - Structural validation
description: Lexer, parser, validator, and ASCII/HTML renderer for UIML design entities.
files:
    - internal/uiml/ast.go
    - internal/uiml/lexer.go
    - internal/uiml/parser.go
    - internal/uiml/render_ascii.go
    - internal/uiml/render_html.go
    - internal/uiml/render_wireframe.go
    - internal/uiml/validate.go
id: COM-0012
kind: component
name: UIML Parser
notes:
    - 'Sync cleanup: render_wireframe.go changes are acknowledged in the UIML Parser component so strict sync can pass before Codex hooks enforce the finish gate.'
purpose: Parse and render UIML (UI Markup Language) design entities
relationships:
    - target: syde-cli
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Lex, parse, validate, and render UIML bodies for design kind entities
slug: uiml-parser-sjdk
updated_at: "2026-04-15T06:06:09Z"
---
