---
category: gotcha
confidence: medium
description: 'fmt.Sprintf eats literal ''%'' chars in CSS — ''width: 80%'' becomes a format directive and silently corrupts the output. When embedding a giant CSS block + body via Sprintf, either escape every % as %% (fragile) or split into wireframeDocPrefix + body + wireframeDocSuffix string concat. The lesson: any large template string with ''%'' literals should NOT go through Sprintf.'
discovered_at: "2026-04-15T03:18:48Z"
entity_refs:
    - uiml-parser
id: LRN-0028
kind: learning
name: 'fmt.Sprintf eats literal ''%'' chars in CSS — ''width: 80%'' b'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: fmtsprintf-eats-literal-chars-in-css-width-80-b-feqg
source: session-observation
updated_at: "2026-04-15T03:18:48Z"
---
