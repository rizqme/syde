---
id: REQ-0278
kind: requirement
name: Install Skill Invocation
slug: install-skill-invocation-2k4w
relationships:
    - target: install-skill-bji4
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:47Z"
statement: When the user runs syde install-skill, the syde CLI shall write the skill files into the .claude directory and append the mandatory syde rules section to CLAUDE.md.
req_type: interface
priority: must
verification: integration test invoking syde install-skill and checking for .claude/skills/syde and CLAUDE.md rules block
source: manual
source_ref: contract:install-skill-bji4
requirement_status: active
rationale: Skill installation is how syde injects its enforcement behavior into Claude Code sessions.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:47Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:47Z"
---
