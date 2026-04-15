---
id: REQ-0165
kind: requirement
name: Skill Installer Writes Claude SKILL Files
slug: skill-installer-writes-claude-skill-files-hek6
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:36Z"
statement: When syde install-skill is invoked for Claude, the skill installer shall write the embedded SKILL.md and references under .claude/skills/syde/.
req_type: functional
priority: must
verification: integration test invoking syde install-skill and inspecting .claude/skills/syde
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Shipping the SKILL content is the installer's primary job.
---
