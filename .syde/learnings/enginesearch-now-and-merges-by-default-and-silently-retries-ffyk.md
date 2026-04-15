---
category: pattern
confidence: medium
description: Engine.Search now AND-merges by default and silently retries with OR when AND yields zero, marking results Broadened=true so the formatter prints a banner. Loose human queries like 'concept entity' or 'relationship label' get useful results without the agent having to remember --any. Strict queries that find AND matches never trigger the fallback.
discovered_at: "2026-04-14T09:36:12Z"
entity_refs:
    - query-engine
id: LRN-0011
kind: learning
name: Engine.Search now AND-merges by default and silently retries
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: enginesearch-now-and-merges-by-default-and-silently-retries-ffyk
source: session-observation
updated_at: "2026-04-14T09:36:12Z"
---
