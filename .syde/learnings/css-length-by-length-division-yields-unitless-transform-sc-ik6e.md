---
category: gotcha
confidence: high
description: 'CSS length-by-length division yields unitless: transform: scale(calc(100cqi / 1440px)) is valid — dividing length by length produces a scalar. This enables fluid scaling of fixed-px layouts inside container-type: inline-size. calc(100cqw / 1440) is INVALID (produces ~0.07cqw length). Used in internal/uiml/render_wireframe.go to scale the 1440x810 wireframe canvas to any container width.'
discovered_at: "2026-04-15T03:38:24Z"
id: LRN-0033
kind: learning
name: 'CSS length-by-length division yields unitless: transform: sc'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: css-length-by-length-division-yields-unitless-transform-sc-ik6e
source: render_wireframe.go implementation
updated_at: "2026-04-15T03:38:24Z"
---
