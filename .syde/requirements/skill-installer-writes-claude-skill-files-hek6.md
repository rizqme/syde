---
id: REQ-0165
kind: requirement
name: Skill Installer Writes Claude SKILL Files
slug: skill-installer-writes-claude-skill-files-hek6
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:50Z"
statement: When syde install-skill is invoked for Claude, the skill installer shall write the embedded SKILL.md and references under .claude/skills/syde/.
req_type: functional
priority: must
verification: integration test invoking syde install-skill and inspecting .claude/skills/syde
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Shipping the SKILL content is the installer's primary job.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:50Z"
---
