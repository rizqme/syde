---
description: How architecture context is auto-loaded into Claude at session start.
edge_cases: No .syde/ exists → skip hook; BadgerDB lock held by another syde process → retry briefly; tree not yet scanned → snapshot notes stale
failure_modes: Corrupted tree.yaml blocks context generation → user runs 'syde tree scan' manually. Stale index with deleted files → user runs 'syde reindex'.
goal: Load full architecture context into the agent before any user task runs
happy_path: Hook runs → syde context emits JSON → Claude Code injects snapshot → agent ready in <1s
id: FLW-0001
kind: flow
name: Session Start Bootstrap
narrative: SessionStart hook from .claude/hooks/syde-hooks.json invokes 'syde context --json'. syde loads the project config, the index, and all entities, formats a snapshot of decisions, components, plans, tasks, and recent learnings, and writes it into the agent's context window. Agent begins the session already knowing the architecture.
performance_notes: Full context load is O(entities). Must stay under 1s for typical projects (<500 entities). Dominant cost is markdown parsing, not BadgerDB.
relationships:
    - target: syde
      type: belongs_to
    - target: skill
      type: involves
    - target: session-context
      type: involves
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: session-start-bootstrap-ng80
trigger: Claude Code session starts in a project with a .syde/ directory
updated_at: "2026-04-14T03:27:02Z"
---
