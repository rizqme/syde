---
id: REQ-0332
kind: requirement
name: 'User request: load syde skill and then can you check remaining task in current'
slug: user-request-load-syde-skill-and-then-can-you-check-remaining-task-in-current-r9vt
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T10:04:51Z"
statement: When the user asks Codex to inspect remaining syde task state, the syde workflow shall load syde context and report the current pending plan tasks.
req_type: functional
priority: must
verification: Manual inspection confirms the response used syde context and listed pending current-plan tasks.
source: user
source_ref: codex:019d93cd-a768-7703-b52c-5b1b7e204988:019d93cd-d822-7f03-8cbc-5f04779b9b0d
requirement_status: active
rationale: The user explicitly asked to load the syde skill and check remaining tasks in the current work.
approved_at: "2026-04-16T01:00:34Z"
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:51Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:51Z"
---
