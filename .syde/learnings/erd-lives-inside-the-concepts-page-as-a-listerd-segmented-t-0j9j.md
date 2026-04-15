---
category: pattern
confidence: medium
description: ERD lives inside the Concepts page as a List/ERD segmented toggle, NOT as a separate sidebar item. conceptView state in App.tsx survives entity selection but resets to 'list' whenever activeKind leaves 'concept'. Clicking a node in ERD navigates to the entity detail AND flips conceptView back to list so the detail panel becomes visible. Keeps the ERD page component reusable; no routing changes needed.
discovered_at: "2026-04-14T10:34:07Z"
entity_refs:
    - web-spa
id: LRN-0018
kind: learning
name: ERD lives inside the Concepts page as a List/ERD segmented t
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: erd-lives-inside-the-concepts-page-as-a-listerd-segmented-t-0j9j
source: session-observation
updated_at: "2026-04-14T10:34:07Z"
---
