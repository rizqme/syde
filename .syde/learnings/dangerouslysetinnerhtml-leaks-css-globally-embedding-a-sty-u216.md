---
category: gotcha
confidence: high
description: 'dangerouslySetInnerHTML leaks CSS globally: embedding a <style> block via dangerouslySetInnerHTML applies its rules to the entire dashboard, not just the injected HTML. Any body/* selector pollutes the parent page. Fix: scope every rule under a unique wrapper class (e.g. .wf-root), wrap output in a matching div.'
discovered_at: "2026-04-15T03:38:24Z"
id: LRN-0035
kind: learning
name: 'dangerouslySetInnerHTML leaks CSS globally: embedding a <sty'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dangerouslysetinnerhtml-leaks-css-globally-embedding-a-sty-u216
source: dashboard wireframe embed
updated_at: "2026-04-15T03:38:24Z"
---
