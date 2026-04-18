---
id: REQ-0100
kind: requirement
name: Skill Requires Hooks Json
slug: skill-requires-hooks-json-bu6o
relationships:
    - target: skill-7fmf
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:50Z"
statement: The syde CLI shall require a hooks_json asset on every installed skill instance.
req_type: constraint
priority: must
verification: integration test confirming .claude/hooks/syde-hooks.json is written by install-skill
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Hooks are required to enforce workflow automation like session bootstrap.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:50Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:50Z"
---
