---
category: gotcha
confidence: medium
description: 'UIML lexer is fragile around tag attributes — even the official ''syde design create'' skeleton (with direction="vertical") triggers parse errors. Root cause: the lexer''s readTextContent slurps everything until the next ''<'', so attribute name=value pairs collapse into a single TokText that the attribute loop can''t decompose. Workaround: stick to attribute-free structural tags (<screen>, <sidebar>, <main>, <heading>, <text>, <button>, <list>, <item>, <card>, <panel>) for any UIML you want to round-trip through the validator. Real fix needs a stateful lexer with attribute-context mode.'
discovered_at: "2026-04-14T11:40:50Z"
entity_refs:
    - uiml-parser
id: LRN-0024
kind: learning
name: UIML lexer is fragile around tag attributes — even the off
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: uiml-lexer-is-fragile-around-tag-attributes-even-the-off-9ay0
source: session-observation
updated_at: "2026-04-14T11:40:50Z"
---
