---
category: pattern
confidence: medium
description: syde wireframe render <slug> [--format html|ascii|image] [--out path] [--open] is the terminal counterpart to the dashboard's wireframe panel. Image mode shells out to headless Google Chrome (auto-detected via wfChromeBinary on macOS/linux) to produce a PNG from a temp HTML file. ASCII mode reuses uiml.RenderASCII. HTML mode pipes the raw renderer output. Use this for visual iteration without booting syded.
discovered_at: "2026-04-15T03:18:48Z"
entity_refs:
    - cli-commands
id: LRN-0030
kind: learning
name: syde wireframe render <slug> [--format html|ascii|image] [--
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: syde-wireframe-render-slug-format-htmlasciiimage-zmuo
source: session-observation
updated_at: "2026-04-15T03:18:48Z"
---
