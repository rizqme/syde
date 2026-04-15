---
id: REQ-0166
kind: requirement
name: Skill Installer Writes Hooks JSON
slug: skill-installer-writes-hooks-json-0vto
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:39Z"
statement: When syde install-skill is invoked, the skill installer shall write .claude/hooks/syde-hooks.json containing the PostToolUse, SessionStart, and Stop hook definitions.
req_type: functional
priority: must
verification: integration test inspecting hooks JSON after install
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Hooks drive the automated syde guardrails inside agent sessions.
---
