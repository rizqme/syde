---
id: REQ-0096
kind: requirement
name: Skill Installer Is Idempotent
slug: skill-installer-is-idempotent-1u2y
relationships:
    - target: skill-7fmf
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:50Z"
statement: When the user runs syde install-skill repeatedly, the syde CLI shall produce the same installed skill state without duplicating CLAUDE.md rule sections.
req_type: functional
priority: must
verification: integration test running syde install-skill twice and diffing the workspace
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Idempotence lets users re-run install without fear of corrupting their project.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:50Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:50Z"
---
