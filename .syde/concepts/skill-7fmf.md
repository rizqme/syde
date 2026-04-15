---
attributes:
    - description: skill identifier (always 'syde' for this project)
      name: name
    - description: the canonical workflow + concept rules + cookbook
      name: SKILL_md
    - description: auxiliary docs (commands.md, entity-spec.md, clarify-guide.md, sync-workflow.md)
      name: references
    - description: PreToolUse / PostToolUse / SessionStart wiring
      name: hooks_json
description: The Claude Code skill (SKILL.md + references + hooks) embedded in the syde binary.
id: CPT-0011
invariants: SKILL.md is the source of truth for agent behavior. Installer is idempotent. Hooks must not modify .syde/ files without going through syde CLI.
kind: concept
lifecycle: Edit skill/SKILL.md → rebuild binary → 'syde install-skill' writes into user project → Claude Code loads on session start
meaning: The Claude Code skill definition — SKILL.md plus references and hooks — that enforces the syde workflow
name: Skill
relationships:
    - target: syde
      type: belongs_to
    - target: skill-installer
      type: references
slug: skill-7fmf
structure_notes: Embedded in the syde binary via go:embed from skill/. Installed into .claude/skills/syde/ + .claude/hooks/syde-hooks.json + CLAUDE.md append rules.
updated_at: "2026-04-14T10:48:03Z"
---
