---
category: gotcha
confidence: medium
description: 'CSS emitted via dangerouslySetInnerHTML into a React component leaks into the parent page globally — rules like ''body { padding: 24px }'' and ''* { box-sizing }'' will apply to the dashboard, not just the embed. Fix: scope EVERY rule under a wrapper class (e.g. .wf-root) and wrap the emit in that class. Browser parser strips <html><body> when the HTML is injected as innerHTML, but style blocks and their rules still execute globally. Bit by this when render_wireframe.go''s CSS was first leaking into the Concepts page.'
discovered_at: "2026-04-15T03:31:55Z"
entity_refs:
    - web-spa
id: LRN-0032
kind: learning
name: CSS emitted via dangerouslySetInnerHTML into a React compone
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: css-emitted-via-dangerouslysetinnerhtml-into-a-react-compone-id76
source: session-observation
updated_at: "2026-04-15T03:31:55Z"
---
