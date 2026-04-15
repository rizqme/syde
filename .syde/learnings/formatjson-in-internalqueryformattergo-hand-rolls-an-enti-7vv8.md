---
category: gotcha
confidence: medium
description: FormatJSON in internal/query/formatter.go hand-rolls an entityMap per kind via type-switch. Every new kind-specific field added to an entity struct MUST be mirrored here or it will silently fail to surface through the dashboard API — even if the data is on disk. Bit by the concept ERD Attributes/Actions not rendering on the Decision panel despite being back-filled correctly; fix was three map assignments in the concept case.
discovered_at: "2026-04-14T10:34:07Z"
entity_refs:
    - query-engine
id: LRN-0016
kind: learning
name: FormatJSON in internal/query/formatter.go hand-rolls an enti
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: formatjson-in-internalqueryformattergo-hand-rolls-an-enti-7vv8
source: session-observation
updated_at: "2026-04-14T10:34:07Z"
---
