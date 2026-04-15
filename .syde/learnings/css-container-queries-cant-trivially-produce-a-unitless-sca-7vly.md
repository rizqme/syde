---
category: pattern
confidence: medium
description: 'CSS container queries can''t trivially produce a unitless scale factor — calc(100cqw / 1440) evaluates to a length (~0.07cqw), not a dimensionless number, so transform: scale() rejects it. The reliable cross-browser pattern for ''scale fixed-size HTML to fit parent width while preserving aspect ratio'' is <svg viewBox="0 0 W H" preserveAspectRatio="xMidYMid meet"><foreignObject width="W" height="H"><div xmlns="http://www.w3.org/1999/xhtml">...HTML...</div></foreignObject></svg>. SVG viewBox math is rock solid and works in Chrome/Firefox without container queries. Used for screen wireframe rendering in render_wireframe.go with 1440×810 (16:9).'
discovered_at: "2026-04-15T03:31:55Z"
entity_refs:
    - uiml-parser
id: LRN-0031
kind: learning
name: CSS container queries can't trivially produce a unitless sca
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: css-container-queries-cant-trivially-produce-a-unitless-sca-7vly
source: session-observation
updated_at: "2026-04-15T03:31:55Z"
---
