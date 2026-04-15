---
acceptance: uiml-survey.md is 2-4 pages covering all four sections; lexer trace shows the exact failing call site.
affected_entities:
    - uiml-parser
completed_at: "2026-04-15T02:44:04Z"
created_at: "2026-04-15T02:37:51Z"
details: Use syde tree context internal/uiml/lexer.go internal/uiml/parser.go internal/uiml/ast.go internal/uiml/render_html.go (NEVER naive Read). Walk the lexer with input '<layout direction="vertical">' showing token-by-token output. Document the inTag-flag minimal patch with line numbers. Confirm parseElement attribute loop is correct. Catalog NodeKinds. For each structural NodeKind, document the Tailwind classes RenderHTML emits today.
id: TSK-0082
kind: task
name: Survey internal/uiml and write uiml-survey.md sections 1-4
objective: .syde/research/uiml-survey.md exists with lexer trace + parser confirmation + AST tag list + per-NodeKind RenderHTML baseline
plan_phase: phase_2
plan_ref: uiml-wireframe-research
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: survey-internaluiml-and-write-uiml-surveymd-sections-1-4-fjmm
task_status: completed
updated_at: "2026-04-15T02:44:04Z"
---
