---
id: COM-0010
kind: component
name: Skill Installer
slug: skill-installer-wbmu
description: Embeds and writes skill files, hooks, and CLAUDE.md rules into target projects
purpose: Ship the SKILL.md + hooks + references from the syde binary into a project's .claude/ directory
notes:
    - SKILL.md extended with bootstrap batch-script pattern
    - ' tree-summarize concurrency caveat'
    - ' and orphan-file resolution playbook (session 2026-04-13).'
    - All references files updated for required-description rule (2026-04-14).
    - Codex hooks template now uses a robust command wrapper that prefers syde on PATH and falls back to ./syde for source checkout workflows.
    - Codex hooks template now uses plain 'syde codex-hook' commands without a ./syde fallback, matching the expected installed syde PATH workflow.
files:
    - internal/skill/hooks.go
    - internal/skill/installer.go
    - internal/skill/templates.go
    - skill/SKILL.md
    - skill/codex/SKILL.md
    - skill/codex/hooks.json
    - skill/embed.go
    - skill/hooks.json
    - skill/references/clarify-guide.md
    - skill/references/commands.md
    - skill/references/entity-spec.md
    - skill/references/requirement-derivation.md
    - skill/references/sync-workflow.md
relationships:
    - target: syde-cli
      type: belongs_to
updated_at: "2026-04-17T11:14:08Z"
responsibility: Render and write skill files including the Finding-aware sync check instructions, hooks.json overlap reminder, and CLAUDE.md append rules
capabilities:
    - Write .claude/skills/syde/SKILL.md + references/
    - Write .claude/hooks/syde-hooks.json (PostToolUse + SessionStart + SessionEnd)
    - Append idempotent syde rules section to CLAUDE.md
boundaries: Does NOT own skill content (that lives in skill/*, embedded via go:embed). Does NOT run hooks.
---
