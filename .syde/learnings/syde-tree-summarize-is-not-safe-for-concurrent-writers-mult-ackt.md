---
category: gotcha
confidence: high
description: 'syde tree summarize is not safe for concurrent writers. Multiple parallel processes (e.g. subagents summarizing different files at once) race on .syde/tree.yaml.tmp rename and some updates are lost. Workaround: subagents must retry failed writes sequentially, or tree summarize should use file locking / atomic CAS.'
discovered_at: "2026-04-13T04:42:29Z"
id: LRN-0001
kind: learning
name: syde tree summarize is not safe for concurrent writers. Mult
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: syde-tree-summarize-is-not-safe-for-concurrent-writers-mult-ackt
source: session-observation
updated_at: 2026-04-13T04:53:19Z
---
