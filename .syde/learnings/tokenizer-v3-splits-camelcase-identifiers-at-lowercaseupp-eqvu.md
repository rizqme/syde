---
category: pattern
confidence: medium
description: Tokenizer v3 splits CamelCase identifiers at lowercase→uppercase boundaries (and snake/dash on the existing alnum split). 'IndexSchemaVersion' indexes as indexschemaversion + index + schema + version, so a search for any sub-word still hits. Acronym-then-lowercase ('URLPath') is intentionally not split — full lowercased form still indexed for exact matches.
discovered_at: "2026-04-14T09:36:12Z"
entity_refs:
    - storage-engine
id: LRN-0012
kind: learning
name: Tokenizer v3 splits CamelCase identifiers at lowercase→upp
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: tokenizer-v3-splits-camelcase-identifiers-at-lowercaseupp-eqvu
source: session-observation
updated_at: "2026-04-14T09:36:12Z"
---
