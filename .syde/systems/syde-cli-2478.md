---
id: SYS-0004
kind: system
name: syde CLI
slug: syde-cli-2478
description: Cobra-based command-line binary 'syde'. Primary interface for all entity CRUD, plans, tasks, tree management, validation, and skill installation.
purpose: Give agents and humans one ergonomic CLI that speaks the full syde model without opening markdown files directly.
relationships:
    - target: syde
      type: belongs_to
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
updated_at: "2026-04-15T07:08:44Z"
context: Invoked on every session start (SessionStart hook), every file write (PostToolUse hook), and throughout a planning/implementation loop.
scope: 'In: all subcommands under ''syde'' (add, get, list, update, query, plan, task, tree, validate, etc.). Out: HTTP serving (syded owns that).'
---
