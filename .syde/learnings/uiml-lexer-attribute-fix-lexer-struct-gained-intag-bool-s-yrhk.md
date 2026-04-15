---
category: pattern
confidence: medium
description: 'UIML lexer attribute fix: Lexer struct gained ''inTag bool'' set true after TokTagOpen, cleared on TokGT/TokSelfClose. When inTag is true Next() reads an identifier via readTagName (returned as TokText) instead of falling through to readTextContent which would slurp everything to the next ''<''. ~10 line patch. The parser was already correct — only the lexer needed the state. <layout direction=''horizontal''>, <grid cols=''4''>, <sidebar width=''200''> all parse cleanly now.'
discovered_at: "2026-04-15T03:18:48Z"
entity_refs:
    - uiml-parser
id: LRN-0029
kind: learning
name: 'UIML lexer attribute fix: Lexer struct gained ''inTag bool'' s'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: uiml-lexer-attribute-fix-lexer-struct-gained-intag-bool-s-yrhk
source: session-observation
updated_at: "2026-04-15T03:18:48Z"
---
