---
id: FLW-0001
kind: flow
name: Session Start Bootstrap
slug: session-start-bootstrap-ng80
description: How architecture context is auto-loaded into Claude at session start.
relationships:
    - target: syde
      type: belongs_to
    - target: skill
      type: involves
    - target: session-context
      type: involves
    - target: architecture-context-shall-load-fast-52uc
      type: references
updated_at: "2026-04-16T10:59:47Z"
trigger: Claude Code session starts in a project with a .syde/ directory
goal: Load full architecture context into the agent before any user task runs
steps:
    - id: s1
      action: SessionStart hook fires
      contract: session-context
      description: Runs syde context
      on_success: s2
    - id: s2
      action: System loads entities and formats snapshot
      description: Reads all entities from store
      on_success: s3
    - id: s3
      action: Agent receives architecture context
      description: Context injected into session
      on_success: done
narrative: SessionStart hook from .claude/hooks/syde-hooks.json invokes 'syde context --json'. syde loads the project config, the index, and all entities, formats a snapshot of decisions, components, plans, and tasks, and writes it into the agent's context window. Agent begins the session already knowing the architecture.
happy_path: Hook runs → syde context emits JSON → Claude Code injects snapshot → agent ready in <1s
edge_cases: No .syde/ exists → skip hook; BadgerDB lock held by another syde process → retry briefly; tree not yet scanned → snapshot notes stale
failure_modes: Corrupted tree.yaml blocks context generation → user runs 'syde tree scan' manually. Stale index with deleted files → user runs 'syde reindex'.
performance_notes: Full context load is O(entities). Must stay under 1s for typical projects (<500 entities). Dominant cost is markdown parsing, not BadgerDB.
---
