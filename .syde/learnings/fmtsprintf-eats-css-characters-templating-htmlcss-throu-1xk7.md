---
category: gotcha
confidence: high
description: 'fmt.Sprintf eats CSS % characters: templating HTML+CSS through fmt.Sprintf breaks on rules like ''width: 80%'' — the % is read as a format directive. Fix: split into prefix/suffix string constants and plain concat, not Sprintf. Bit the render_wireframe.go build.'
discovered_at: "2026-04-15T03:38:24Z"
id: LRN-0034
kind: learning
name: 'fmt.Sprintf eats CSS % characters: templating HTML+CSS throu'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: fmtsprintf-eats-css-characters-templating-htmlcss-throu-1xk7
source: render_wireframe.go debugging
updated_at: "2026-04-15T03:38:24Z"
---
