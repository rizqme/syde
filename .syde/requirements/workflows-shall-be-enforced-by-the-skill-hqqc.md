---
id: REQ-0012
kind: requirement
name: Workflows shall be enforced by the skill
slug: workflows-shall-be-enforced-by-the-skill-hqqc
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:51Z"
statement: The syde install-skill command shall write hooks that block session-end and code modifications when the syde model is out of sync.
req_type: constraint
priority: must
verification: Inspect installed .claude/hooks/syde-hooks.json for SessionStart/Stop/PostToolUse entries
source: manual
source_ref: system:syde-5tdt:design_principles
requirement_status: active
rationale: Without enforcement the model rots; the skill is the only mechanism keeping agents in the loop.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:51Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:51Z"
---
