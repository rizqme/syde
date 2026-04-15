---
category: pattern
confidence: medium
description: 'UIML lexer attribute bug root cause: lexer.Next() falls through to readTextContent (which slurps everything until next ''<'') when inside an opening tag, so ''direction="vertical">'' becomes one giant TokText that the parser cannot decompose. Minimal fix: add inTag bool flag to Lexer, set true after TokTagOpen and clear on TokGT/TokSelfClose; when inTag is true, read identifier via readTagName instead of falling through to readTextContent. ~10 lines. Parser''s attribute loop already handles distinct identifier/equals/value tokens correctly. Documented in detail in .syde/research/uiml-survey.md section 1.'
discovered_at: "2026-04-15T02:51:06Z"
entity_refs:
    - uiml-parser
id: LRN-0027
kind: learning
name: 'UIML lexer attribute bug root cause: lexer.Next() falls thro'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: uiml-lexer-attribute-bug-root-cause-lexernext-falls-thro-28ft
source: session-observation
updated_at: "2026-04-15T02:51:06Z"
---
