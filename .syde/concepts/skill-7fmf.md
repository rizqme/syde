---
id: CPT-0011
kind: concept
name: Skill
slug: skill-7fmf
description: The Claude Code skill (SKILL.md + references + hooks) embedded in the syde binary.
relationships:
    - target: syde
      type: belongs_to
    - target: skill-installer
      type: implemented_by
updated_at: "2026-04-17T10:30:16Z"
meaning: The Claude Code skill definition — SKILL.md plus references and hooks — that enforces the syde workflow
lifecycle: Edit skill/SKILL.md → rebuild binary → 'syde install-skill' writes into user project → Claude Code loads on session start
invariants: SKILL.md is the source of truth for agent behavior. Installer is idempotent. Hooks must not modify .syde/ files without going through syde CLI.
---
