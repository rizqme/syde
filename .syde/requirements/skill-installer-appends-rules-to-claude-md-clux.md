---
id: REQ-0167
kind: requirement
name: Skill Installer Appends Rules To CLAUDE md
slug: skill-installer-appends-rules-to-claude-md-clux
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:49Z"
statement: When syde install-skill is invoked, the skill installer shall append an idempotent syde rules section to the project CLAUDE.md file.
req_type: functional
priority: must
verification: integration test that repeated installs leave CLAUDE.md stable
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Idempotent append keeps the rules current without duplicating content.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:49Z"
---
