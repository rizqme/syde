---
id: REQ-0100
kind: requirement
name: Skill Requires Hooks Json
slug: skill-requires-hooks-json-bu6o
relationships:
    - target: skill-7fmf
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:15Z"
statement: The syde CLI shall require a hooks_json asset on every installed skill instance.
req_type: constraint
priority: must
verification: integration test confirming .claude/hooks/syde-hooks.json is written by install-skill
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Hooks are required to enforce workflow automation like session bootstrap.
---
