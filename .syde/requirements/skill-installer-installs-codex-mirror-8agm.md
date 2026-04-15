---
id: REQ-0168
kind: requirement
name: Skill Installer Installs Codex Mirror
slug: skill-installer-installs-codex-mirror-8agm
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:44Z"
statement: Where Codex support is requested, the skill installer shall mirror the embedded skill under .agents/skills/syde and update .codex and AGENTS.md idempotently.
req_type: functional
priority: must
verification: integration test invoking syde install-skill --codex
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Codex and Claude users share the same syde skill content.
---
