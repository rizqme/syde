---
category: pattern
confidence: medium
description: scripts/wireframe-shot.sh + headless Google Chrome (/Applications/Google Chrome.app/Contents/MacOS/Google Chrome --headless --disable-gpu --screenshot) is the inspection loop for any dashboard-rendering work. Auto-launches syded if dead, resolves project slug from /api/projects by matching cwd, dumps PNG to /tmp/wireframe-<slug>.png. The Read tool can display the PNG so an agent can SEE the live render and iterate visually. Used during the uiml-wireframe-research plan to validate the wireframe sandbox against user reference images.
discovered_at: "2026-04-15T02:51:06Z"
entity_refs:
    - web-spa
id: LRN-0026
kind: learning
name: scripts/wireframe-shot.sh + headless Google Chrome (/Applica
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: scriptswireframe-shotsh-headless-google-chrome-applica-qwwx
source: session-observation
updated_at: "2026-04-15T02:51:06Z"
---
