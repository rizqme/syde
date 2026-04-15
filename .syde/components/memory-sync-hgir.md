---
boundaries: Does NOT modify entities. Does NOT call LLMs.
capabilities:
    - Write per-entity memory files under .claude/
    - List and clean stale memory entries
description: Generator for Claude Code per-entity memory files under .claude/.
files:
    - internal/memory/manager.go
id: COM-0008
kind: component
name: Memory Sync
purpose: Keep Claude Code memory files in sync with the current syde entity state
relationships:
    - target: syde-cli
      type: belongs_to
    - target: storage-engine
      type: depends_on
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Generate .claude/ memory files summarizing entities for agent session context
slug: memory-sync-hgir
updated_at: "2026-04-14T03:35:54Z"
---
