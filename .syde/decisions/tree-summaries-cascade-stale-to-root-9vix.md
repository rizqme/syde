---
alternatives_considered: Only mark the changed file stale (ancestors rot silently); snapshot-only tree with no staleness tracking (no guarantee of freshness)
category: data
consequences: Every session that touches source must end with 'syde tree status --strict' exit 0. SessionEnd hook gates on this. Subagent delegation is mandatory for file summarization to keep main-context token usage down.
description: Every file change cascades stale up to root; sessions can't end with stale nodes.
id: DEC-0005
kind: decision
name: Tree Summaries Cascade Stale To Root
rationale: Folder summaries are derived from their children. If a child changes, the folder's summary is by definition out of date. Cascading stale up ensures no one forgets to re-summarize the parent. The leaves-first loop with subagent delegation keeps the main context cheap.
relationships:
    - target: syde
      type: applies_to
    - target: summary-tree
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: tree-summaries-cascade-stale-to-root-9vix
statement: Any file whose hash changes marks itself stale plus every ancestor folder up to the root. Any summary update marks the direct parent stale. Sessions must not end with stale nodes.
tradeoffs: More stale noise on large changes. Mitigated by --leaves-only filtering.
updated_at: "2026-04-14T03:27:03Z"
---
