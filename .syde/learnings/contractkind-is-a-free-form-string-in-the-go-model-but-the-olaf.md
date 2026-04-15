---
category: pattern
confidence: medium
description: 'contract_kind is a free-form string in the Go model but the skill recommends an enum. When modelling schema-style contracts (KV key prefix, SQL table, proto, cache key, queue payload), use contract_kind=storage and interaction_pattern=schema. Existing 7 BadgerDB key contracts were mis-classified as event/pub-sub and are now corrected. The dashboard filter bar was extended with contract_kind: and pattern: DSL keys that pull dynamic values from loaded entities.'
discovered_at: "2026-04-14T11:13:21Z"
entity_refs:
    - skill-installer
id: LRN-0023
kind: learning
name: 'contract_kind is a free-form string in the Go model but the '
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: contractkind-is-a-free-form-string-in-the-go-model-but-the-olaf
source: session-observation
updated_at: "2026-04-14T11:13:21Z"
---
