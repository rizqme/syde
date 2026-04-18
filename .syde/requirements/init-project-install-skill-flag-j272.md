---
id: REQ-0277
kind: requirement
name: Init Project Install Skill Flag
slug: init-project-install-skill-flag-j272
relationships:
    - target: init-project-xm6c
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:46Z"
statement: Where --install-skill is passed to syde init, the syde CLI shall also install the Claude Code skill files after initializing the project.
req_type: interface
priority: must
verification: integration test invoking syde init --install-skill
source: manual
source_ref: contract:init-project-xm6c
requirement_status: active
rationale: Combining init and skill install reduces setup friction for new projects.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:46Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:46Z"
---
