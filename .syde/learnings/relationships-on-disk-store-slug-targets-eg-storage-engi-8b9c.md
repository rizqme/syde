---
category: gotcha
confidence: medium
description: 'Relationships on disk store slug targets (e.g. ''storage-engine''), not entity IDs. GetInbound(ID) misses them all. Use the resolver''s alias trick: query GetInbound with every form (ID, full slug, bare slug) and de-dupe. oneHopRelated in ByFile does this now.'
discovered_at: "2026-04-14T08:29:48Z"
entity_refs:
    - query-engine
id: LRN-0006
kind: learning
name: Relationships on disk store slug targets (e.g. 'storage-engi
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: relationships-on-disk-store-slug-targets-eg-storage-engi-8b9c
source: session-observation
updated_at: "2026-04-14T08:29:48Z"
---
