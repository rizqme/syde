---
boundaries: Does NOT call LLMs — summaries must be written by an agent or human via CLI. Does NOT persist entities.
capabilities:
    - SHA-256 hashing with cascade-stale on change
    - gitignore + built-in ignore pattern matching
    - Binary/large-file auto-summary
    - Ancestor breadcrumb + node summary + inlined file content bundle
    - Ignored-node flag for orphan-file validator escape hatch
description: Hashed file/folder mirror with cascade-stale change tracking and context bundles.
files:
    - internal/tree/model.go
    - internal/tree/store.go
    - internal/tree/walk.go
    - internal/tree/scan.go
    - internal/tree/summarize.go
    - internal/tree/context.go
    - internal/tree/render.go
    - internal/tree/ignore.go
id: COM-0005
kind: component
name: Summary Tree
notes:
    - 'Codex compatibility: summary tree default ignore list now excludes generated .codex/ and .agents/ directories, matching existing .claude/ treatment.'
    - 'Codex compatibility: summary tree default ignores now exclude generated agent integration files and directories (.codex/, .agents/, AGENTS.md, CLAUDE.md).'
purpose: Track every source file/folder with a stored summary and detect drift against the code
relationships:
    - target: syde-cli
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Walk, hash, diff, summarize, and render the project file tree
slug: summary-tree-fq6u
updated_at: "2026-04-15T06:05:46Z"
---
