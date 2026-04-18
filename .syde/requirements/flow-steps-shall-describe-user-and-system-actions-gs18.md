---
id: REQ-0349
kind: requirement
name: Flow steps shall describe user and system actions
slug: flow-steps-shall-describe-user-and-system-actions-gs18
description: Step action text names the actor
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:46Z"
statement: When a flow has structured steps, the syde skill shall require each step action to describe what the user or system does at that point in the journey.
req_type: usability
priority: should
verification: Skill docs teach action text conventions
source: plan
requirement_status: active
rationale: Steps like 'do thing' are useless; 'User runs syde plan create' is actionable
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:46Z"
---
