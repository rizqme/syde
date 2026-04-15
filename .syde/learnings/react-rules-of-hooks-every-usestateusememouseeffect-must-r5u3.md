---
category: gotcha
confidence: high
description: 'React rules-of-hooks: every useState/useMemo/useEffect must be called in the same order on every render. If you early-return on loading/error state BEFORE calling a later hook, React throws error #310 as soon as the loading flag flips. Fix: move all hook calls above all early returns, and guard the hook bodies on nullable inputs instead.'
discovered_at: "2026-04-13T14:07:23Z"
entity_refs:
    - web-spa
id: LRN-0003
kind: learning
name: 'React rules-of-hooks: every useState/useMemo/useEffect must '
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: react-rules-of-hooks-every-usestateusememouseeffect-must-r5u3
source: session-observation
updated_at: "2026-04-13T14:07:23Z"
---
