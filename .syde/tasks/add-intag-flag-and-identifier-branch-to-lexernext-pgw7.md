---
acceptance: syde design create Foo && syde design show foo produces zero parse warnings. Creating a screen contract with --wireframe '<screen direction="horizontal"><sidebar width="200"><heading>Sidebar</heading></sidebar><main><heading>Main</heading></main></screen>' passes syde sync check.
affected_entities:
    - uiml-parser
affected_files:
    - internal/uiml/lexer.go
completed_at: "2026-04-15T03:00:16Z"
created_at: "2026-04-15T02:53:51Z"
details: 'internal/uiml/lexer.go: add inTag bool to Lexer struct. In Next(): on TokTagOpen branch set inTag=true; on TokGT branch clear inTag=false; on TokSelfClose branch clear inTag=false. Add a new branch before readTextContent: if inTag is true, read an identifier via readTagName and return as TokText. The existing TokEquals and readQuotedString branches already work.'
id: TSK-0086
kind: task
name: Add inTag flag and identifier branch to Lexer.Next
objective: Lexer tokenizes attributes correctly when inside an opening tag
plan_phase: phase_1
plan_ref: uiml-wireframe-render
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-intag-flag-and-identifier-branch-to-lexernext-pgw7
task_status: completed
updated_at: "2026-04-15T03:00:16Z"
---
