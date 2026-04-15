---
category: gotcha
confidence: medium
description: 'syde remove <plan-slug> reported ''Removed plan: <name>'' but the .md file was still on disk after the command exited. The sync check later surfaced the phantom plan as still-present. Root cause not yet investigated — add to the CLI remove command audit. Workaround: manually rm the file and run syde reindex.'
discovered_at: "2026-04-14T09:06:18Z"
entity_refs:
    - cli-commands
id: LRN-0009
kind: learning
name: 'syde remove <plan-slug> reported ''Removed plan: <name>'' but '
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: syde-remove-plan-slug-reported-removed-plan-name-but-mrfj
source: session-observation
updated_at: "2026-04-14T09:06:18Z"
---
