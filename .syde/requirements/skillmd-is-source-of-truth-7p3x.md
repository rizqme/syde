---
id: REQ-0094
kind: requirement
name: SKILL.md Is Source Of Truth
slug: skillmd-is-source-of-truth-7p3x
relationships:
    - target: skill-7fmf
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:50Z"
statement: The syde CLI shall treat the embedded skill/SKILL.md as the sole source of truth for skill behavior installed into user projects.
req_type: constraint
priority: must
verification: code review confirming install-skill reads from the embedded skill asset and not external locations
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: A single source prevents drift between the binary and installed skill files.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:50Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:50Z"
---
