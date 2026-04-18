---
id: REQ-0098
kind: requirement
name: Hooks Must Not Bypass CLI For Syde Mutations
slug: hooks-must-not-bypass-cli-for-syde-mutations-pvrb
relationships:
    - target: skill-7fmf
      type: refines
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:46Z"
statement: The skill hooks shall not modify any file under .syde/ except through the syde CLI.
req_type: constraint
priority: must
verification: code review of hooks.json and hook scripts
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Going through the CLI keeps the index, tree, and file-store consistent.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:46Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:46Z"
---
