---
id: REQ-0168
kind: requirement
name: Skill Installer Installs Codex Mirror
slug: skill-installer-installs-codex-mirror-8agm
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:49Z"
statement: Where Codex support is requested, the skill installer shall mirror the embedded skill under .agents/skills/syde and update .codex and AGENTS.md idempotently.
req_type: functional
priority: must
verification: integration test invoking syde install-skill --codex
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Codex and Claude users share the same syde skill content.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:49Z"
---
